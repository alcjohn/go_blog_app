package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string
	Password  string
}
