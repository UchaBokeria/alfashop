package posts

import (
	"github.com/labstack/echo/v4"

	"main/internal/types/dto"
	"main/pkg/engine/controller"
)

func Register(app *echo.Group) {
	app.GET("/posts", controller.Set[Filters](index))
	app.GET("/posts/:id", controller.Set[dto.ByID](detail))
}
