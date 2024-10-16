package categories

import (
	"main/internal/types/dto"
	"main/pkg/engine/controller"

	"github.com/labstack/echo/v4"
)

func Register(admin *echo.Group) {
	categories := admin.Group("/categories")
	categories.GET("", controller.Set[any](index))
	categories.GET("/create/:id", controller.Set[dto.ByID](createForm))
	categories.GET("/update/:id", controller.Set[dto.ByID](updateForm))

	categories.POST("", controller.Set[CategoriesDto](create))
	categories.PUT("/:id", controller.Set[CategoriesDto](update))
	categories.DELETE("/:id", controller.Set[dto.ByID](remove))
	categories.PUT("/:status/:id", controller.Set[dto.Visibility](status))
}
