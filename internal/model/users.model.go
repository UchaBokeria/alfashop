package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Fullname		string
	Email			string
	Password		string
	Token			string
	// `gorm:"constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
}