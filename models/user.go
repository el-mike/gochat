package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
