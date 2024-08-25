package chat

import (
	"github.com/labstack/echo/v4"

	"main/pkg/engine/controller"
)

func Register(app *echo.Group) {
	Chat := app.Group("/chat")
	Chat.POST("", controller.Set[MailStrategyDto](index))
	Chat.POST("/mailStrategy", controller.Set[MailStrategyDto](MailStrategy))
	// Chat.GET("/new", controller.Set[any](NewChat))
	// Chat.GET("/:id", controller.Set[any](OldChat))
	// Chat.GET(":id", controller.Set[any](index))
}