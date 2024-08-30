package categories

import (
	"main/internal/model"
	"main/pkg/helpers"
	"main/pkg/storage"
)

var Seed = []model.Categories{
	{
		ParentID: nil,
		Level:    1,
		Name:     "Root",
		Slug:     "root",
		Public:   true,
		IconID:   15,
		Filters:  CategoryFilters[0:1],
	},
	{
		ParentID: nil,
		Level:    1,
		Name:     "Root 2",
		Slug:     "root_2",
		Public:   true,
		IconID:   15,
		Filters:  CategoryFilters[0:1],
	},
	{
		ParentID: helpers.GetPointedInt(1),
		Level:    2,
		Name:     "middle",
		Slug:     "middle",
		Public:   true,
		IconID:   15,
		Filters:  CategoryFilters[0:1],
	},
	{
		ParentID:        helpers.GetPointedInt(3),
		Level:           3,
		Name:            "baby",
		Slug:            "baby",
		Description:     "Baby",
		DescriptionHTML: "<p>Baby</p>",
		Public:          true,
		IconID:          15,
		Filters:         nil,
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)
	}
}
