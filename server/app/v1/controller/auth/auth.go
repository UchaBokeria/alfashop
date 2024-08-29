package auth

import (
	"github.com/labstack/echo/v4"

	"main/pkg/engine/controller"
)

func Register(app *echo.Group) {
	auth := app.Group("/auth")
	auth.GET("/login", controller.Set[LoginDto](Login))
	auth.POST("/signup", controller.Set[SignupDto](Signup))
	auth.POST("/verifyPhone", controller.Set[VerifyPhoneDto](VerifyPhone))
	auth.POST("/verifyEmail", controller.Set[VerifyEmailDto](VerifyEmail))
	auth.POST("/forgotPassword", controller.Set[ForgotPasswordDto](ForgotPassword))
	auth.POST("/resetPassword", controller.Set[ResetPasswordDto](ResetPassword))
	auth.POST("/changePassword", controller.Set[ChangePasswordDto](ChangePassword))
	auth.POST("/logout", controller.Set[any](Logout))
}
