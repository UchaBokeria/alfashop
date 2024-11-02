package cart

import (
	"main/internal/enums"
	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/helpers"
	"main/pkg/storage"
	"main/server/admin/repository"
	view "main/server/app/v1/view/pages/orders"
	"net/http"
	"strings"
	"time"
)

func refresh(ctx *controller.Context, Param *CartDto) error {
	var products []model.Products

	storage.DB.
		Where("ID in ?", func() []uint {
			var filters []uint
			ids := strings.Split(Param.IDs, ",")
			for _, id := range ids {
				filters = append(filters, helpers.Uint(id))
			}
			return filters
		}()).
		Where(model.Products{
			Public: true,
		}).
		Preload("Pics").
		Preload("Category").
		Preload("Properties").
		Find(&products)

	return ctx.JSON(http.StatusOK, products)
}

func checkout(ctx *controller.Context, Param *CheckoutDto) error {
	Total := 0.00
	Products := make([]model.Products, 0)

	if ctx.User().ID == 0 {
		return ctx.Redirect(http.StatusMovedPermanently, "/auth/signup")
	}

	for i := 0; i < len(Param.Items); i++ {
		var product model.Products
		storage.DB.Find(&product, uint(Param.Items[i].ID))

		rawPrice := float64(Param.Items[i].Quantity) * product.Price
		price := rawPrice - (rawPrice * product.Discount)
		Total += price

		Products = append(Products, product)
	}

	Order := model.Orders{
		Total:    Total,
		Products: Products,
		User:     ctx.User(),
		Deadline: time.Now(),
		Comment:  Param.Comment,
		StatusID: enums.OrderStatus.New,
	}

	if err := repository.Create[model.Orders](&Order); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.Html(view.Successful(Order))
}
