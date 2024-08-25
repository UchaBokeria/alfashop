package branches

import (
	"github.com/labstack/echo/v4"

	"main/internal/types/dto"
	"main/pkg/engine/controller"
)

func Register(admin *echo.Group) {
	branches := admin.Group("/branches")
	branches.GET("", controller.Set[any](index))
	branches.GET("/:id", controller.Set[dto.ByID](unique))
	branches.POST("", controller.Set[CreateBranchesDto](create))
	branches.PUT("/:id", controller.Set[UpdateBranchesDto](update))
	branches.DELETE("/:id", controller.Set[dto.ByID](remove))

	branches.GET("/properties", controller.Set[any](properties))
	branches.GET("/districts/:id/:indicator", controller.Set[DistrictsDto](districts))
	
}
