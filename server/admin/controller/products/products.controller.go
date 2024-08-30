package products

import (
	"main/internal/model"
	"main/internal/types/dto"
	"main/pkg/engine/controller"
	"main/pkg/helpers"
	"main/pkg/storage"
	"main/server/admin/repository"
	"main/server/admin/view/pages/products"

	"github.com/a-h/templ"
)

func index(ctx *controller.Context) error {
	var productPages []model.Products
	storage.DB.Preload("Pics").Preload("Category").Find(&productPages)
	return ctx.Html(products.Page(productPages))
}

func unique(ctx *controller.Context, ID *dto.ByID) error {
	var Product model.Products
	var Categories []model.Categories

	storage.DB.Preload("Category").First(&Product, ID)

	subQuery := storage.DB.Model(&model.Categories{}).
		Select("DISTINCT parent_id").
		Where("parent_id IS NOT NULL")

	storage.DB.Order("parent_id ASC").
		Where("id NOT IN (?)", subQuery).
		Find(&Categories, &model.Categories{Public: true})

	CategoriesSelect, _ := helpers.ToSelect[model.Categories](Categories, "ID", "Name")
	return ctx.Html(products.Update(Product, CategoriesSelect))
}

func newForm(ctx *controller.Context) error {
	return ctx.Html(templ.NopComponent)
}

// func create(ctx *controller.Context, params *ProductsDto) error {
// 	repository.Create[model.Products](&model.Products{
// 		Name: params.Name,
// 		Slug: params.Slug,
// 		Public: params.Public,
// 		CategoryID: params.CategoryID,
// 		Description: params.Description,
// 		DescriptionHtml: params.DescriptionHtml,
// 		Pics: params.Pics,
// 		TechnicalSheet: params.TechnicalSheet,
// 	})

// 	return index(ctx)
// }

func update(ctx *controller.Context, params *ProductsDto) error {

	repository.Update[model.Products](params.ID, &model.Products{
		Name:            params.Name,
		Public:          params.Public,
		Description:     params.Description,
		DescriptionHtml: params.DescriptionHtml,
	})

	return index(ctx)
}

func remove(ctx *controller.Context, ID *dto.ByID) error {
	repository.Delete[model.Products](ID)
	return index(ctx)
}
