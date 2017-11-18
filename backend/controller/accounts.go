package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"repo.letscode.sii.pl/wroclaw/three/backend/middleware"
	"repo.letscode.sii.pl/wroclaw/three/backend/model"
	"repo.letscode.sii.pl/wroclaw/three/backend/utils"
)

type jwtResponse struct {
	Token string `json:"token"`
}

type Accounts struct {
	Database *gorm.DB
}

func (a *Accounts) Register(router *mux.Router) {
	router.HandleFunc("/register/", a.HandleRegister).Methods(http.MethodPost)
	router.HandleFunc("/login/", a.HandleLogin).Methods(http.MethodPost)
	router.HandleFunc("/token/", a.HandleRefresh).Methods(http.MethodPost)
}

func (a *Accounts) generateJWT(user *model.User) (string, error) {
	// generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, middleware.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		User: user,
	})
	tok, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return tok, fmt.Errorf("error occured while generating JWT: %s", err)
	}
	return tok, nil
}

func (a *Accounts) HandleRegister(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	user := model.User{}
	if err := decoder.Decode(&user); err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, model.ErrAccountsUnknown).AppendDebug(err).Write(rw)
		return
	}

	if err := user.Validate(); err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, err).Write(rw)
		return
	}

	// check if user already exists
	var existingUser model.User

	res := a.Database.First(&existingUser, "email = ?", user.Email)

	if !res.RecordNotFound() {
		utils.NewErrorResponse(http.StatusBadRequest, model.ErrUserEmailRegistered).Write(rw)
		return
	}

	// encrypt password
	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, model.ErrAccountsUnknown).AppendDebug(err).Write(rw)
		return
	}

	user.Password = string(pwd)

	if err := a.Database.Create(&user).Error; err != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, model.ErrAccountsUnknown).AppendDebug(err).Write(rw)
		return
	}

	// clear the password for struct reuse
	user.Password = ""

	tok, err := a.generateJWT(&user)

	if err != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, model.ErrAccountsUnknown).AppendDebug(err).Write(rw)
		return
	}

	// return JWT to client
	body, _ := json.Marshal(&jwtResponse{
		Token: tok,
	})
	rw.WriteHeader(http.StatusOK)
	rw.Write(body)
}

func (a *Accounts) HandleLogin(rw http.ResponseWriter, r *http.Request) {
	// decode request
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	user := model.User{}
	if err := decoder.Decode(&user); err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, model.ErrAccountsParsingError).AppendDebug(err).Write(rw)
		return
	}

	if err := user.Validate(); err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, err).Write(rw)
		return
	}

	var dbUser model.User
	res := a.Database.First(&dbUser, "email = ?", user.Email)

	if res.RecordNotFound() {
		utils.NewErrorResponse(http.StatusNotFound, model.ErrUserEmailNotFound).Write(rw)
		return
	}

	if res.Error != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, model.ErrUserEmailNotFound).AppendDebug(res.Error).Write(rw)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, model.ErrUserPasswordInvalid).Write(rw)
		return
	}

	dbUser.Password = ""

	// generate JWT
	tok, err := a.generateJWT(&dbUser)
	if err != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, err).Write(rw)
		return
	}

	// return JWT to client
	body, _ := json.Marshal(&jwtResponse{
		Token: tok,
	})
	rw.WriteHeader(http.StatusOK)
	rw.Write(body)
}

func (a *Accounts) HandleRefresh(rw http.ResponseWriter, r *http.Request) {
	tok, claims, err := middleware.ParseToken(r)
	if err != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrAccountsParsingError.Error()},
			DebugErrors: []string{err.Error()},
			Code:        http.StatusInternalServerError,
		}).Write(rw)
		return
	}

	if ve, ok := err.(*jwt.ValidationError); !tok.Valid && (!ok || ve.Errors&jwt.ValidationErrorExpired == 0) {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrAccountsParsingError.Error()},
			DebugErrors: []string{err.Error()},
			Code:        http.StatusBadRequest,
		}).Write(rw)
		return
	}

	user := model.User{}
	if res := a.Database.First(&user, claims.User.ID); res.Error != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrAccountsUnknown.Error()},
			DebugErrors: []string{fmt.Sprintf("error occured while querying the database: %s", res.Error.Error())},
			Code:        http.StatusInternalServerError,
		}).Write(rw)
		return
	}

	// generate JWT
	token, err := a.generateJWT(&user)
	if err != nil {
		(&utils.ErrorResponse{
			Errors: []string{err.Error()},
			Code:   http.StatusInternalServerError,
		}).Write(rw)
		return
	}

	// return JWT to client
	body, _ := json.Marshal(&jwtResponse{
		Token: token,
	})
	rw.WriteHeader(http.StatusOK)
	rw.Write(body)
}
