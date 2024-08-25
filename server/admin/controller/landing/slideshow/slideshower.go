package slideshow

import (
	"main/internal/types/dto"
	"main/pkg/engine/controller"

	"github.com/labstack/echo/v4"
)

func Register(setting *echo.Group) {
	slideshow := setting.Group("/slideshow")
	slideshow.GET("", controller.Set[any](index))
	slideshow.POST("", controller.Set[SlideshowDto](SlideshowNew))
	slideshow.POST("/:id", controller.Set[SlideshowDto](Slideshow))
	slideshow.DELETE("/:id", controller.Set[dto.ByID](SlideRemove))
}
