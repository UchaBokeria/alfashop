package auth

import (
	"main/pkg/engine/controller"

	"github.com/labstack/echo/v4"
)

func Register(echo *echo.Echo) {
	auth := echo.Group("/auth")
	auth.GET("/login", controller.Set[LoginDto](login))
	auth.POST("/login", controller.Set[LoginDto](login))
	// auth.POST("/register", controller.Set[any](register))
	// auth.PUT("/changepwd", controller.Set[LoginDto](changePassword))
}