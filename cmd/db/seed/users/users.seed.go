package users

import (
	"main/internal/model"
	"main/pkg/helpers"
	"main/pkg/storage"
	"time"
)

var Seed = []model.Users{
	{
		Fullname:        "Administrator",
		Email:           "admin@alfashop.ge",
		Password:        "$2a$12$BI/hpA2GqPofEtlZnmInF.NMIRGket6YnAF8cJ0XwIkQLPWGI7dpK",
		Token:           "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
		EmailVerified:   true,
		Phone:           "599 156 862",
		EmailVerifyedAt: time.Now(),
		PhoneVerified:   true,
		PhoneVerifiedAt: time.Now(),
		TypeID:          1,
		Company:         helpers.ToPointer("AlfaShop"),
		CityID:          helpers.ToPointer(1),
	},
}

var Types = []model.UserTypes{
	{
		Name: "admin",
	},
	{
		Name: "client",
	},
	{
		Name: "dealer",
	},
}

func Populate() {
	for _, row := range Types {
		storage.DB.Create(&row)
	}
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
