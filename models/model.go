package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"username"`
	Password string `json:"password"`
}
