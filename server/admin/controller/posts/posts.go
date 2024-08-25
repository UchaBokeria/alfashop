package posts

import (
	"github.com/labstack/echo/v4"

	"main/internal/types/dto"
	"main/pkg/engine/controller"
)

func Register(setting *echo.Group) {
	post := setting.Group("/posts")
	post.GET("", controller.Set[any](index))
	post.POST("", controller.Set[PostsDto](create))
	post.POST("/:id", controller.Set[PostsDto](update))
	post.DELETE("/:id", controller.Set[dto.ByID](remove))
}
