package settings

import (
	"github.com/labstack/echo/v4"

	"main/pkg/engine/controller"
	"main/server/admin/controller/landing/about"
	"main/server/admin/controller/landing/contact"
	"main/server/admin/controller/landing/events"
	"main/server/admin/controller/landing/faq"
	newvideos "main/server/admin/controller/landing/newVideos"
	"main/server/admin/controller/landing/news"
	"main/server/admin/controller/landing/presentation"
	"main/server/admin/controller/landing/products"
	"main/server/admin/controller/landing/slideshow"
	"main/server/admin/controller/landing/socials"
)

func Register(admin *echo.Group) {
	admin.GET("", controller.Set[Param](index))
	admin.GET("/setting/:tab", controller.Set[Param](index))

	settings := admin.Group("/settings")
	slideshow.Register(settings)
	presentation.Register(settings)
	products.Register(settings)
	news.Register(settings)
	events.Register(settings)
	contact.Register(settings)
	socials.Register(settings)
	faq.Register(settings)
	socials.Register(settings)
	about.Register(settings)
	newvideos.Register(settings)
}
