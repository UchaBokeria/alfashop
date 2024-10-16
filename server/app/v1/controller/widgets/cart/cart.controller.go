package cart

import (
	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/helpers"
	"main/pkg/storage"
	"net/http"
	"strings"
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
	return ctx.HTML(http.StatusOK, "widgets/cart/checkout")
}
