package model

import (
	"errors"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"repo.letscode.sii.pl/wroclaw/three/backend/utils"
)

var (
	ErrUserEmailInvalid     = errors.New("email invalid")
	ErrUserPasswordInvalid  = errors.New("password invalid")
	ErrUserNameInvalid      = errors.New("name invalid")
	ErrUserEmailRegistered  = errors.New("mail already registered")
	ErrUserEmailNotFound    = errors.New("email not found")
	ErrAccountsUnknown      = errors.New("unknown error occured")
	ErrAccountsParsingError = errors.New("token parsing error occured")
)

type User struct {
	ID        uint      `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email,omitempty" gorm:"index"`
	Password  string    `json:"password,omitempty"`
	Projects  []Project `json:"projects,omitempty" gorm:"many2many:user_projects;"`
}

func (u *User) Validate() error {
	// validate input
	errors := []error{}
	if len(u.Email) == 0 {
		errors = append(errors, ErrUserEmailInvalid)
	}
	if len(u.Password) == 0 {
		errors = append(errors, ErrUserPasswordInvalid)
	}
	if len(u.Name) == 0 {
		errors = append(errors, ErrUserNameInvalid)
	}
	if len(errors) > 0 {
		return utils.NewErrorResponse(http.StatusBadRequest, errors...)
	}
	return nil
}

func (u *User) AddToProject(db *gorm.DB, project *Project) error {
	assoc := db.Model(u).Association("Projects")
	if assoc := assoc.Append(project); assoc.Error != nil {
		return assoc.Error
	}
	return nil
}
