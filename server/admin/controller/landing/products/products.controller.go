package products

import (
	"main/internal/model"
	"main/internal/types/dto"
	"main/pkg/engine/controller"
	"main/pkg/helpers"
	"main/pkg/storage"
	"main/server/admin/repository"
	"main/server/admin/view/pages/landing/tabs"
)

func index(ctx *controller.Context) error {
	var Interface model.Interface
	var Products []model.Products

	storage.DB.Find(&Products)
	storage.DB.Preload("Products.Pics").
			   Preload("Products.Category").
			   Last(&Interface)

	ProductsSelect, _ := helpers.ToSelect(Products, "ID", "Name")
	return ctx.Html(tabs.Products(Interface.Products, ProductsSelect))
}

func create(ctx *controller.Context, params *productsDto) error {
	storage.DB.Model(repository.FindByID[model.Interface](1)).
			   Association("Products").
			   Append(repository.FindByID[model.Products](params.ProductID))

	return index(ctx)
}

func remove(ctx *controller.Context, ID *dto.ByID) error {
	storage.DB.Model(repository.FindByID[model.Interface](1)).
			   Association("Products").Delete(repository.FindByID[model.Products](ID.ID))

	return index(ctx)
}
