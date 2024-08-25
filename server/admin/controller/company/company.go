package company

import (
	"github.com/labstack/echo/v4"

	"main/internal/types/dto"
	"main/pkg/engine/controller"
)

func Register(admin *echo.Group) {
	company := admin.Group("/company")
	company.GET("", controller.Set[any](index))
	company.POST("", controller.Set[CompanyCreateDto](create))
	company.PUT("/:id", controller.Set[CompanyUpdateDto](update))
	company.DELETE("/:id", controller.Set[dto.ByID](remove))

	company.GET("/create", controller.Set[any](newForm))
	company.GET("/:id", controller.Set[dto.ByID](unique))
}