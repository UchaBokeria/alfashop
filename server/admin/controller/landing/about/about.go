package about

import (
	"main/pkg/engine/controller"

	"github.com/labstack/echo/v4"
)

func Register(setting *echo.Group) {
	setting.GET("/about", controller.Set[any](index))
	setting.POST("/about", controller.Set[AboutDto](contact))
}