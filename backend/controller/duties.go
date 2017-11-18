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

type Duties struct {
	Database *gorm.DB
}

func (d *Duties) Register(router *mux.Router) {
}

func (d *Duties) HandleAdd(rw http.ResponseWriter, r *http.Request) {
	project := r.Context().Value(middleware.ContextProjectKey).(*model.Project)
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	duty := &model.Duty{}
	if err := decoder.Decode(duty); err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, model.ErrDutyInternal).AppendDebug(err).Write(rw)
		return
	}

	if err := duty.Validate(); err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, model.ErrDutyNameInvalid).AppendDebug(err).Write(rw)
		return
	}

	if err := project.AddDuty(d.Database, duty); err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, model.ErrDutyInternal).AppendDebug(err).Write(rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	if byt, err := json.Marshal(duty); err == nil {
		rw.Write(byt)
	}
}

func (d *Duties) HandleConfirm(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, model.ErrProjectIDInvalid).AppendDebug(err).Write(rw)
		return
	}

	duty := &model.Duty{}
	duty.ID = uint(id)

	assoc := d.Database.Model(duty).Association("Confirmations")

	//TODO: security
	if res := assoc.Append(&model.Confirmation{UserID: uint(id)}); res.Error != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, model.ErrProjectInternal).AppendDebug(res.Error).Write(rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (d *Duties) HandleDelete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.NewErrorResponse(http.StatusBadRequest, model.ErrProjectIDInvalid).AppendDebug(err).Write(rw)
		return
	}

	//TODO: security
	if res := d.Database.Model(&model.Duty{}).Where("id = ?", id).Delete(&model.Duty{}); res.Error != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, model.ErrProjectInternal).AppendDebug(res.Error).Write(rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (d *Duties) HandleGetAll(rw http.ResponseWriter, r *http.Request) {
	project := r.Context().Value(middleware.ContextProjectKey).(*model.Project)
	duties := []model.Duty{}
	if res := d.Database.Model(project).Association("Duties").Find(&duties); res.Error != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrProjectInternal.Error()},
			DebugErrors: []string{res.Error.Error()},
			Code:        http.StatusInternalServerError,
		}).Write(rw)
		return
	}

	byt, err := json.Marshal(&duties)
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

func (d *Duties) HandleGetSingle(rw http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value(middleware.ContextUserKey).(*model.User)
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		(&utils.ErrorResponse{
			Errors:      []string{model.ErrProjectIDInvalid.Error()},
			DebugErrors: []string{err.Error()},
			Code:        http.StatusBadRequest,
		}).Write(rw)
		return
	}

	duty := &model.Duty{}
	duty.ID = uint(id)
	if err := duty.Find(d.Database); err != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, err).Write(rw)
		return
	}

	//TODO: security
	if err := d.Database.Model(duty).Related(&duty.Confirmations); err.Error != nil {
		utils.NewErrorResponse(http.StatusInternalServerError, err.Error).Write(rw)
		return
	}

	byt, err := json.Marshal(&duty)
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
