package products

import (
	"main/internal/model"
	"main/pkg/helpers"
	"main/pkg/storage"
	"main/server/admin/repository"
)

var Seed = []model.Products{
	{
		Name:            "Product 1",
		Slug:            "product-1",
		Description:     "Product 1 description",
		DescriptionHtml: `<p>Description</p>`,
		Public:          true,
		Price:           100,
		Pics: []model.Files{
			{
				Name:       `686-e90-384-thumb__600_0_0_0_crop.jpg`,
				Original:   `1 L | 1118114-001`,
				Location:   "local",
				Path:       "/uploads/products/images/686-e90-384-thumb__600_0_0_0_crop.jpg",
				Size:       664,
				Base64:     "",
				Compressed: false,
				TypeID:     helpers.GetDbTypeIdByExtension(`jpg`),
			},
			{
				Name:       `012-88a-943-thumb__600_0_0_0_crop.jpg`,
				Original:   `5 L | 1118114-005`,
				Location:   "local",
				Path:       "/uploads/products/images/012-88a-943-thumb__600_0_0_0_crop.jpg",
				Size:       664,
				Base64:     "",
				Compressed: false,
				TypeID:     helpers.GetDbTypeIdByExtension(`jpg`),
			},
		},
	},
}

func Populate() {
	for _, row := range Seed {
		storage.DB.Create(&row)

		storage.DB.
			Model(&row).
			Association("Category").
			Append(
				repository.FindByID[model.Categories](1),
				repository.FindByID[model.Categories](2),
			)

	}
}
