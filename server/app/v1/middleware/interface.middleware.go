package middleware

import (
	"github.com/labstack/echo/v4"

	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/storage"
	"main/server/app/v1/view/shared"
)

func Interface() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return controller.Use(func(ctx *controller.Context) error {
			if ctx.IsHtmx() {
				return next(ctx)
			}

			var Interface model.Interface
			result := storage.DB.
				Preload("Contact").
				Preload("SocialMedia").
				Preload("Languages.Icon").
				Last(&Interface)
			if result.Error != nil {
				return ctx.Html(shared.ErrorPage())
			}

			ctx.Set("Interface", Interface)
			return next(ctx)
		})
	}
}
