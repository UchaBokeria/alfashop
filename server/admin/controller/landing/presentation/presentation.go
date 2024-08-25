package presentation

import (
	"main/pkg/engine/controller"

	"github.com/labstack/echo/v4"
)

func Register(setting *echo.Group) {
	setting.GET("/presentation", controller.Set[any](index))
	setting.POST("/presentation", controller.Set[PresentationDto](presentation))
}