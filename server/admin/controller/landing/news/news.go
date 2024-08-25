package news

import (
	"github.com/labstack/echo/v4"

	"main/internal/types/dto"
	"main/pkg/engine/controller"
)

func Register(setting *echo.Group) {
	news := setting.Group("/news")
	news.GET("", controller.Set[any](index))
	news.POST("", controller.Set[NewsDto](create))
	news.DELETE("/:id", controller.Set[dto.ByID](remove))
}
