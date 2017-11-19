package model

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	"repo.letscode.sii.pl/wroclaw/three/backend/utils"
)

var (
	ErrProjectInternal    = errors.New("internal error")
	ErrProjectIDInvalid   = errors.New("invalid id")
	ErrProjectNameInvalid = errors.New("invalid name")
)

type Project struct {
	gorm.Model
	Name   string `json:"name"`
	Hidden bool   `json:"hidden"`
	Icon   string `json:"icon,omitempty"`
	Duties []Duty `gorm:"ForeignKey:ProjectID" json:"duties"`
	Users  []User `gorm:"many2many:user_projects;" json:"users"`
}

func (o *Project) Find(db *gorm.DB) error {
	if res := db.First(o, o.ID); res.Error != nil {
		if res.RecordNotFound() {
			return utils.NewErrorResponse(http.StatusNotFound, res.Error)
		}
		return res.Error
	}
	return nil
}

func (o *Project) Add(db *gorm.DB) error {
	errors := []error{}
	if len(o.Name) == 0 {
		errors = append(errors, ErrProjectNameInvalid)
	}

	if len(errors) > 0 {
		return utils.NewErrorResponse(http.StatusBadRequest, errors...)
	}

	if res := db.Set("gorm:save_associations", false).Create(&o); res.Error != nil {
		return &utils.ErrorResponse{
			Errors:      []string{ErrProjectInternal.Error()},
			DebugErrors: []string{res.Error.Error()},
		}
	}

	for _, user := range o.Users {
		if err := user.AddToProject(db, o); err != nil {
			return err
		}
	}

	return nil
}

func (p *Project) AddDuty(db *gorm.DB, d *Duty) error {
	if assoc := db.Model(p).Association("Duties").Append(d); assoc.Error != nil {
		return assoc.Error
	}
	return nil
}
