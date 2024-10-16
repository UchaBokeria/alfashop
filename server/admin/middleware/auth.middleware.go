package middleware

import (
	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/storage"
	admin "main/server/admin/view/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHeaderCreds struct {
	Email string `header:"X-User"`
	Token string `header:"X-Token"`
}

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return controller.Use(func(ctx *controller.Context) error {
			var User model.Users
			var Parameters AuthHeaderCreds

			ctx.Set("ISADMIN", false)
			err := (&echo.DefaultBinder{}).BindHeaders(ctx, &Parameters)
			if err != nil || Parameters.Email == "" || Parameters.Token == "" {
				Parameters.Email = ctx.ReadCookie("user").Value
				Parameters.Token = ctx.ReadCookie("token").Value
			}

			if Parameters.Email == "" || Parameters.Token == "" {
				return ctx.Renders(http.StatusBadRequest, admin.Login())
			}

			result := storage.DB.Where(&model.Users{Email: Parameters.Email, Token: Parameters.Token}).Last(&User)
			if result.Error != nil || result.RowsAffected < 1 {
				return ctx.Renders(http.StatusUnauthorized, admin.Login())
			}

			ctx.Set("ISADMIN", true)
			ctx.Set("USER", User)

			return next(ctx)
		})
	}
}
