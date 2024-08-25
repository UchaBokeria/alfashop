package newvideos

import (
	"main/internal/types/dto"
	"main/pkg/engine/controller"

	"github.com/labstack/echo/v4"
)

func Register(setting *echo.Group) {
	newVideos := setting.Group("/newvideos")
	newVideos.GET("", controller.Set[any](index))
	newVideos.POST("", controller.Set[NewVideoDto](create))
	newVideos.DELETE("/:id", controller.Set[dto.ByID](remove))
}
