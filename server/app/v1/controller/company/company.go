package company

import (
	"github.com/labstack/echo/v4"

	"main/pkg/engine/controller"
)

func Register(app *echo.Group) {
	company := app.Group("/company")
	company.GET("/:page", controller.Set[CompanyDto](index))
	company.GET("/static/faq", controller.Set[any](faq))
}
