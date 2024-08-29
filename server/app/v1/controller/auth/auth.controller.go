package auth

import (
	"main/pkg/engine/controller"
	view "main/server/app/v1/view/pages/auth"
	"net/http"
)

func Login(ctx *controller.Context, Parameters *LoginDto) error {
	return ctx.Html(view.Login())
}

func Signup(ctx *controller.Context, Parameters *SignupDto) error {
	return ctx.Html(view.Signup())
}

func VerifyPhone(ctx *controller.Context, Parameters *VerifyPhoneDto) error {
	return ctx.Html(view.VerifyPhone())
}

func VerifyEmail(ctx *controller.Context, Parameters *VerifyEmailDto) error {
	return ctx.Html(view.VerifyEmail())
}

func ForgotPassword(ctx *controller.Context, Parameters *ForgotPasswordDto) error {
	return ctx.Html(view.ForgotPassword())
}

func ResetPassword(ctx *controller.Context, Parameters *ResetPasswordDto) error {
	return ctx.Html(view.ResetPassword())
}

func ChangePassword(ctx *controller.Context, Parameters *ChangePasswordDto) error {
	return ctx.Html(view.ChangePassword())
}

func Logout(ctx *controller.Context, Parameters *any) error {
	return ctx.String(http.StatusOK, "Logout")
}
