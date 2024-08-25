package app

import (
	v1 "main/server/app/v1"

	"github.com/labstack/echo/v4"
)

func Run(echo *echo.Echo) *echo.Echo {
	v1.Run(echo.Group(""))
	return echo
}
