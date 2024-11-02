package payments

import (
	"main/internal/model"
	"main/pkg/storage"
)

var Seed = []model.PaymentStatus{
	{
		Name: "აქტიური",
	},
	{
		Name: "დასრულებული",
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
