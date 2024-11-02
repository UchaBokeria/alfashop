package orders

import (
	"main/internal/model"
	"main/pkg/storage"
	"main/server/admin/repository"
	"time"
)

var Seed = []model.Orders{
	{
		UserID:   1,
		StatusID: 1,
		Total:    0,
		Deadline: time.Now(),
		Comment:  "",
	},
	{
		UserID:   2,
		StatusID: 2,
		Total:    0,
		Deadline: time.Now(),
		Comment:  "",
	},
}

var status = []model.Order_status{
	{
		Name: "new",
	},
	{
		Name: "pending",
	},
	{
		Name: "inprocess",
	},
	{
		Name: "issued",
	},
	{
		Name: "canceled",
	},
	{
		Name: "done",
	},
}

func Populate() {
	for _, row := range status {
		storage.DB.Create(&row)
	}

	for _, row := range Seed {
		storage.DB.Create(&row)

		storage.DB.
			Model(&row).
			Association("Products").
			Append(
				repository.FindByID[model.Products](1),
			)
	}
}
