package posts

import (
	"main/internal/model"
	"main/pkg/storage"
)

var Type = []model.Posts_types{
	{
		Name: "ფოტო",
		Slug: "photo",
	},
	{
		Name: "ვიდეო",
		Slug: "video",
	},
	{
		Name: "ბლოგი",
		Slug: "blog",
	},
	{
		Name: "ავტოსპორტი",
		Slug: "autosport",
	},
	{
		Name: "ივენთი",
		Slug: "event",
	},
} 

func Types() {
	for _, row := range Type { storage.DB.Create(&row) }
}