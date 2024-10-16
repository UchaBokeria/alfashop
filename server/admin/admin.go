package admin

import (
	"main/server/admin/controller/auth"
	"main/server/admin/controller/branches"
	"main/server/admin/controller/categories"
	"main/server/admin/controller/company"
	settings "main/server/admin/controller/landing"
	"main/server/admin/controller/posts"
	"main/server/admin/controller/products"
	"main/server/admin/middleware"

	"github.com/labstack/echo/v4"
)

func Run(echo *echo.Echo) *echo.Group {
	auth.Register(echo)

	Admin := echo.Group("/admin", middleware.Auth())
	categories.Register(Admin)
	branches.Register(Admin)
	settings.Register(Admin)
	products.Register(Admin)
	company.Register(Admin)
	posts.Register(Admin)
	return Admin
}
