package model

import (
	"errors"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"repo.letscode.sii.pl/wroclaw/three/backend/utils"
)

var (
	ErrDutyInternal      = errors.New("internal error")
	ErrDutyIDInvalid     = errors.New("invalid id")
	ErrDutyNameInvalid   = errors.New("invalid name")
	ErrDutyRewardInvalid = errors.New("invalid duty")
)

type Duty struct {
	ID            uint           `json:"id"`
	CreatedAt     time.Time      `json:",omitempty"`
	ProjectID     uint           `json:"project_id,omitempty"`
	Name          string         `json:"name,omitempty"`
	Reward        int            `json:"reward,omitempty"`
	Confirmations []Confirmation `json:"confirmations"`
}

func (o *Duty) Find(db *gorm.DB) error {
	if res := db.First(o, o.ID); res.Error != nil {
		if res.RecordNotFound() {
			return utils.NewErrorResponse(http.StatusNotFound, res.Error)
		}
		return res.Error
	}
	return nil
}

func (d *Duty) Validate() error {
	errors := []error{}
	if len(d.Name) == 0 {
		errors = append(errors, ErrDutyNameInvalid)
	}

	if d.Reward == 0 {
		errors = append(errors, ErrDutyRewardInvalid)
	}

	if len(errors) > 0 {
		return utils.NewErrorResponse(http.StatusBadRequest, errors...)
	}
	return nil
}

func (d *Duty) Add(db *gorm.DB) error {
	if err := d.Validate(); err != nil {
		return err
	}
	if res := db.Create(&d); res.Error != nil {
		return &utils.ErrorResponse{
			Errors:      []string{ErrProjectInternal.Error()},
			DebugErrors: []string{res.Error.Error()},
		}
	}
	return nil
}
