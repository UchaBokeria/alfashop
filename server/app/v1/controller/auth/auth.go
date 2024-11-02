package auth

import (
	"github.com/labstack/echo/v4"

	"main/pkg/engine/controller"
)

func Register(app *echo.Group) {
	auth := app.Group("/auth")

	auth.GET("/login", controller.Set[any](LoginPage))
	auth.GET("/signup", controller.Set[any](SignupPage))
	auth.GET("/password", controller.Set[any](ForgotPage))

	auth.GET("/verify/email/:email/:token", controller.Set[VerifyEmailDto](VerifyEmail))
	auth.GET("/verify/password/:email/:token", controller.Set[VerifyEmailDto](ForgotPassword))

	auth.POST("/login", controller.Set[LoginDto](Login))
	auth.POST("/signup", controller.Set[SignupDto](Signup))
	auth.POST("/password", controller.Set[ChangePasswordDto](ChangePassword))

	auth.GET("/logout", controller.Set[any](Logout))
}
