package products

import (
	"github.com/labstack/echo/v4"

	"main/pkg/engine/controller"
)

func Register(app *echo.Group) {
	products := app.Group("/products")
	products.GET("", controller.Set[FiltersQuery](index))
	products.GET("/list", controller.Set[FiltersQuery](list))
	products.GET("/:id", controller.Set[ProductDetail](detail))
	products.GET("/categories", controller.Set[any](categories))
}