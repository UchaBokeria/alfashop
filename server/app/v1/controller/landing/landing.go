package landing

import (
	"github.com/labstack/echo/v4"

	"main/pkg/engine/controller"
)

func Register(app *echo.Group) {
	app.GET("/", controller.Set[any](index))
	app.POST("/subscribe", controller.Set[any](subscribe))
}