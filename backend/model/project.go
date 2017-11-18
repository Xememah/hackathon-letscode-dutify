package model

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name   string
	Hidden bool
	Duties []Duty `gorm:"ForeignKey:ProjectID"`
	Users  []User `gorm:"many2many:user_projects;"`
}
