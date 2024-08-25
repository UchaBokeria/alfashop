package faq

import (
	"github.com/labstack/echo/v4"

	"main/internal/types/dto"
	"main/pkg/engine/controller"
)

func Register(setting *echo.Group) {
	faq := setting.Group("/faq")
	faq.POST("", controller.Set[FaqDto](create))
	faq.POST("/:id", controller.Set[FaqDto](update))
	faq.DELETE("/:id", controller.Set[dto.ByID](remove))
}