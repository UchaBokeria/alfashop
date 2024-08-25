package model

import "gorm.io/gorm"

type Faq struct {
	gorm.Model
	Name     string
	Slug     string
	Answer   string
	Question string
}
