package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Username string
	Password string
}
