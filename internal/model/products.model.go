package model

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Name 				string
	Slug 				string
	Description 		string
	DescriptionHtml     string
	Public              bool               		`gorm:"default:true"`
	CategoryID          int
	Category            Categories         		`gorm:"foreignKey:CategoryID"`
	TechnicalSheetID    *int						
	TechnicalSheet      Files					`gorm:"foreignKey:TechnicalSheetID;constraint: OnUpdate:CASCADE, OnDelete:SET NULL;"`
	Pics 				[]Files					`gorm:"many2many:product_pics;"`
	Properties         	[]Product_properties
}

type Product_properties struct {
	gorm.Model
	ProductsID      	uint
	Characteristics		string
	Unit 				string
	Data				string
	Examination			string
}
