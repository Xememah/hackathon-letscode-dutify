package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"repo.letscode.sii.pl/wroclaw/three/backend/middleware"
	"repo.letscode.sii.pl/wroclaw/three/backend/model"
	"repo.letscode.sii.pl/wroclaw/three/backend/utils"
)

type Projects struct {
	Database *gorm.DB
}

func (p *Projects) Register(router *mux.Router) {
	router.Handle("/", middleware.RequiresAuth(http.HandlerFunc(p.HandleGetAll))).Methods(http.MethodGet)
	router.Handle("/", middleware.RequiresAuth(http.HandlerFunc(p.HandleAdd))).Methods(http.MethodPost)

	d := &Duties{Database: p.Database}
	router.Handle("/{projectid:[0-9]+}/duties", middleware.RequiresAuth(middleware.Project(p.Database, http.HandlerFunc(d.HandleGetAll)))).Methods(http.MethodGet)
	router.Handle("/{projectid:[0-9]+}/duties", middleware.RequiresAuth(middleware.Project(p.Database, http.HandlerFunc(d.HandleAdd)))).Methods(http.MethodPost)
	router.Handle("/{projectid:[0-9]+}/duties/{id:[0-9]+}", middleware.RequiresAuth(middleware.Project(p.Database, http.HandlerFunc(d.HandleGetSingle)))).Methods(http.MethodGet)
	router.Handle("/{projectid:[0-9]+}/duties/{id:[0-9]+}", middleware.RequiresAuth(middleware.Project(p.Database, http.HandlerFunc(d.HandleDelete)))).Methods(http.MethodDelete)
	router.Handle("/{projectid:[0-9]+}/duties/{id:[0-9]+}/confirm", middleware.RequiresAuth(middleware.Project(p.Database, http.HandlerFunc(d.HandleConfirm)))).Methods(http.MethodPost)

	router.Handle("/{projectid:[0-9]+}/", middleware.RequiresAuth(http.HandlerFunc(p.HandleGetSingle))).Methods(http.MethodGet)
	router.Handle("/{projectid:[0-9]+}/join", middleware.RequiresAuth(http.HandlerFunc(p.HandleJoin))).Methods(http.MethodPost)
	router.Handle("/{projectid:[0-9]+}/", middleware.RequiresAuth(http.HandlerFunc(p.HandlePutSingle))).Methods(http.MethodPut)
	router.Handle("/{projectid:[0-9]+}/", middleware.RequiresAuth(http.HandlerFunc(p.HandleDelete))).Methods(http.MethodDelete)
}

func (p *Projects) HandleAdd(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(*model.User)
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	project := model.Project{}
	if err := decoder.Decode(&project); err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, err).Write(rw)
		return
	}

	project.Users = append(project.Users, *user)
	if err := project.Add(p.Database); err != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, err).Write(rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	if byt, err := json.Marshal(project); err == nil {
		rw.Write(byt)
	}
}

func (p *Projects) HandleJoin(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(*model.User)
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["projectid"])
	if err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, model.ErrProjectIDInvalid).AppendDebug(err).Write(rw)
		return
	}

	project := &model.Project{}
	project.ID = uint(id)
	if err := project.Find(p.Database); err != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, err).Write(rw)
		return
	}

	if err := user.AddToProject(p.Database, project); err != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, err).Write(rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (p *Projects) HandleDelete(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(*model.User)
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["projectid"])
	if err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, model.ErrProjectIDInvalid).AppendDebug(err).Write(rw)
		return
	}

	//TODO: security

	if res := p.Database.Model(user).Where("id = ?", id).Delete(&model.Project{}); res.Error != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrProjectInternal.Error()},
			DebugErrors: []string{res.Error.Error()},
			Code:        http.StatusInternalServerError,
		}).Write(rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (p *Projects) HandleGetAll(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUserKey).(*model.User)
	projects := []model.Project{}
	res := p.Database
	if res := res.Model(&user).Association("Projects").Find(&projects); res.Error != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrProjectInternal.Error()},
			DebugErrors: []string{res.Error.Error()},
		}).Write(rw)
		return
	}

	byt, err := json.Marshal(&projects)
	if err != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrProjectInternal.Error()},
			DebugErrors: []string{err.Error()},
			Code:        http.StatusInternalServerError,
		}).Write(rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(byt)
}

func (p *Projects) HandleGetSingle(rw http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value(middleware.ContextUserKey).(*model.User)
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["projectid"])
	if err != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrProjectIDInvalid.Error()},
			DebugErrors: []string{err.Error()},
			Code:        http.StatusBadRequest,
		}).Write(rw)
		return
	}

	project := &model.Project{}
	project.ID = uint(id)
	res := p.Database

	//TODO: security

	if res := res.First(project, uint(id)); res.Error != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrProjectInternal.Error()},
			DebugErrors: []string{res.Error.Error()},
			Code:        http.StatusInternalServerError,
		}).Write(rw)
		return
	}

	if res := res.Model(project).Related(&project.Duties); res.Error != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, model.ErrProjectInternal).AppendDebug(res.Error).Write(rw)
		return
	}

	if res := res.Select("id, name").Model(project).Related(&project.Users, "Users"); res.Error != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, model.ErrProjectInternal).AppendDebug(res.Error).Write(rw)
		return
	}

	type Result struct {
		ID     uint `json:"user" db:"id"`
		Points int  `json:"points" db:"points"`
	}
	results := []Result{}
	if res := p.Database.Table("confirmations").Select("confirmations.user_id as id, sum(reward) as points").Joins("INNER JOIN duties ON duties.id = confirmations.duty_id AND duties.project_id=?", vars["projectid"]).Group("confirmations.user_id").Find(&results); res.Error != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, model.ErrProjectInternal).AppendDebug(res.Error).Write(rw)
		return
	}

	type Resp struct {
		*model.Project
		Ranking []Result `json:"ranking"`
	}

	byt, err := json.Marshal(&Resp{
		Project: project,
		Ranking: results,
	})
	if err != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrProjectInternal.Error()},
			DebugErrors: []string{err.Error()},
			Code:        http.StatusInternalServerError,
		}).Write(rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(byt)
}

func (p *Projects) HandlePutSingle(rw http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value(middleware.ContextUserKey).(*model.User)
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["projectid"])
	if err != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrProjectIDInvalid.Error()},
			DebugErrors: []string{err.Error()},
			Code:        http.StatusBadRequest,
		}).Write(rw)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	project := model.Project{}
	if err := decoder.Decode(&project); err != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrProjectInternal.Error()},
			DebugErrors: []string{err.Error()},
			Code:        http.StatusBadRequest,
		}).Write(rw)
		return
	}

	project.ID = uint(id)
	db := p.Database

	// TODO: security

	if res := db.Set("gorm:save_associations", false).Save(&project); res.Error != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrProjectInternal.Error()},
			DebugErrors: []string{res.Error.Error()},
			Code:        http.StatusInternalServerError,
		}).Write(rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
