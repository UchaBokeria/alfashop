package events

import (
	"github.com/labstack/echo/v4"

	"main/internal/types/dto"
	"main/pkg/engine/controller"
)

func Register(setting *echo.Group) {
	events := setting.Group("/events")
	events.GET("", controller.Set[any](index))
	events.POST("", controller.Set[EventsDto](create))
	events.DELETE("/:id", controller.Set[dto.ByID](remove))
}
