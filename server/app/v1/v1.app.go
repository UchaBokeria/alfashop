package v1

import (
	"github.com/labstack/echo/v4"

	"main/server/app/v1/controller/branches"
	"main/server/app/v1/controller/company"
	"main/server/app/v1/controller/landing"
	"main/server/app/v1/controller/posts"
	"main/server/app/v1/controller/products"
	"main/server/app/v1/controller/widgets/chat"
	"main/server/app/v1/middleware"
)

func Run(app *echo.Group) {
	/* Middlewares */
	app.Use(middleware.Interface())
	// app.Use(middleware.Upload())

	/* Widgets */
	chat.Register(app)

	/* Pages */
	landing.Register(app)
	products.Register(app)
	branches.Register(app)
	posts.Register(app)
	company.Register(app)
}
