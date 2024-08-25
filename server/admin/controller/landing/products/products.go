package products

import (
	"main/internal/types/dto"
	"main/pkg/engine/controller"

	"github.com/labstack/echo/v4"
)

func Register(setting *echo.Group) {
	products := setting.Group("/products")
	products.GET("", controller.Set[any](index))
	products.POST("", controller.Set[productsDto](create))
	products.DELETE("/:id", controller.Set[dto.ByID](remove))
}
