package products

import (
	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/helpers"
	"main/pkg/storage"
	view "main/server/app/v1/view/pages/products"
	"net/http"
	"strconv"
)

func index(ctx *controller.Context, Filters *FiltersQuery) error {
	var Category model.Categories
	var Categories []model.Categories

	ID := Filters.Category

	subQuery := storage.DB.Model(&model.Categories{}).
		Select("DISTINCT parent_id").
		Where("parent_id IS NOT NULL")

	storage.DB.Order("parent_id ASC").Where("id NOT IN (?)", subQuery).Find(&Categories, &model.Categories{Public: true})
	storage.DB.Where(&model.Categories{Public: true}).First(&Category, ID)

	return ctx.Html(view.ProductLayout(Category, Categories))
}

func categories(ctx *controller.Context) error {
	var Categories []model.Categories
	storage.DB.Preload("Icon").Where(&model.Categories{Public: true}).Find(&Categories)
	return ctx.Html(view.ProductCategories(Categories))
}

// func Paginate(ctx *controller.Context) func(db *gorm.DB) *gorm.DB {
// 	return func (db *gorm.DB) *gorm.DB {
// 		pageSize := ctx.PageSize()
// 		offset := (ctx.Page() - 1) * pageSize
// 		return db.Offset(offset).Limit(pageSize)
// 	}
// }

func list(ctx *controller.Context, Filters *FiltersQuery) error {
	var Products []model.Products
	var Where model.Products = model.Products{Public: true}
	pageSize := ctx.PageSize()
	offset := (ctx.Page() - 1) * pageSize

	query := storage.DB.Offset(offset).Limit(pageSize)

	if Filters.Category != 0 {
		CategoryID := Filters.Category
		Where.CategoryID = CategoryID
		query = query.Where(Where)
	}

	if Filters.Searcher != "" {
		query.Where("name ILIKE ?", "%"+Filters.Searcher+"%")
	}

	NextPage := strconv.Itoa(ctx.Page() + 1)

	query.Preload("Pics").
		Preload("Category").
		Preload("Properties").
		Find(&Products)

	if len(Products) == 0 {
		return ctx.String(http.StatusOK, "")
	}
	return ctx.Html(view.ProductList(NextPage, Filters.Searcher, helpers.String(Filters.Category), Products))
}

func detail(ctx *controller.Context, Filters *ProductDetail) error {
	var Product model.Products
	ID := Filters.ID

	storage.DB.Preload("Pics").
		Preload("Category").
		Preload("Properties").
		Find(&Product, ID)

	return ctx.Html(view.ProductDetail(Product))
}
