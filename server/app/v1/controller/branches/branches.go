package branches

import (
	"github.com/labstack/echo/v4"

	"main/pkg/engine/controller"
)

func Register(app *echo.Group) {
	app.GET("/branches", controller.Set[BranchesDto](index))
}