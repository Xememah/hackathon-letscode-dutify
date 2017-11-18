package model

import "github.com/jinzhu/gorm"

type Duty struct {
	gorm.Model
	ProjectID     uint
	Name          string
	Reward        int
	Confirmations []Confirmation
}
