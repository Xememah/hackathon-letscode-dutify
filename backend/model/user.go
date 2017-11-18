package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string   `json:"name,omitempty"`
	Email    string   `json:"email" gorm:"index"`
	Password string   `json:"password,omitempty"`
}
