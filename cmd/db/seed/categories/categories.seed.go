package categories

import (
	"main/internal/model"
	"main/pkg/helpers"
	"main/pkg/storage"

	"gorm.io/gorm"
)

var Seed = []model.Categories{
	{
		ParentID: nil,
		Level:    1,
		Name:     "Root",
		Slug:     "root",
		Public:   true,
		IconID:   helpers.GetPointedInt(15),
		Filters:  CategoryFilters[0:1],
		Children: []*model.Categories{
			{
				Level:   2,
				Name:    "middle",
				Slug:    "middle",
				Public:  true,
				IconID:  helpers.GetPointedInt(15),
				Filters: CategoryFilters[0:1],
				Children: []*model.Categories{
					{
						Level:           3,
						Name:            "baby",
						Slug:            "baby",
						Description:     "Baby",
						DescriptionHTML: "<p>Baby</p>",
						Public:          true,
						IconID:          helpers.GetPointedInt(15),
						Filters:         nil,
					},
				},
			},
		},
	},
	{
		ParentID: nil,
		Level:    1,
		Name:     "Root 2",
		Slug:     "root_2",
		Public:   true,
		IconID:   helpers.GetPointedInt(15),
		Filters:  CategoryFilters[0:1],
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&row)
	}
}
