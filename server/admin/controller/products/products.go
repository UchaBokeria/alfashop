package products

import (
	"github.com/labstack/echo/v4"

	"main/internal/types/dto"
	"main/pkg/engine/controller"
)

func Register(admin *echo.Group) {
	product := admin.Group("/products")
	product.GET("", controller.Set[any](index))
	// product.POST("", controller.Set[ProductsDto](create))
	product.PUT("/:id", controller.Set[ProductsDto](update))
	product.DELETE("/:id", controller.Set[dto.ByID](remove))

	product.GET("/create", controller.Set[any](newForm))
	product.GET("/:id", controller.Set[dto.ByID](unique))
}