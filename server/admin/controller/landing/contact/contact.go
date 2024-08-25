package contact

import (
	"main/pkg/engine/controller"

	"github.com/labstack/echo/v4"
)

func Register(setting *echo.Group) {
	setting.GET("/contacts", controller.Set[any](index))
	setting.POST("/contacts", controller.Set[ContactDto](contact))
}