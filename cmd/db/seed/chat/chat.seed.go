package chat

import (
	"main/internal/model"
	"main/pkg/storage"
)

var ChatStatus []model.Chat_status = []model.Chat_status{
	{
		Name: "აქტიური",
		Slug: "open",
	},
	{
		Name: "დასრულებული",
		Slug: "closed",
	},
}


var ChatType []model.Chat_status = []model.Chat_status{
	{
		Name: "ლაივი",
		Slug: "live",
	},
	{
		Name: "ელფოსტა",
		Slug: "mail",
	},
}

func Populate() {
	for _, row := range ChatType { storage.DB.Create(&row) }
	for _, row := range ChatStatus { storage.DB.Create(&row) }
}