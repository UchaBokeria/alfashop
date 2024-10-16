package cart

import (
	"github.com/labstack/echo/v4"

	"main/pkg/engine/controller"
)

func Register(app *echo.Group) {
	Cart := app.Group("/cart")
	Cart.POST("", controller.Set[CheckoutDto](checkout))
	Cart.GET("/refresh", controller.Set[CartDto](refresh))
}
