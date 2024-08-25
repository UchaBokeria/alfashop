package socials

import (
	"main/pkg/engine/controller"

	"github.com/labstack/echo/v4"
)

func Register(setting *echo.Group) {
	setting.GET("/socials", controller.Set[any](index))
	setting.POST("/socials", controller.Set[SocialDto](Socials))
}