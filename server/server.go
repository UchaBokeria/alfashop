package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"main/pkg/engine/controller"
	lang "main/pkg/engine/language"
	"main/pkg/globals"
	"main/pkg/storage"
	"main/server/admin"
	"main/server/app"
)

type (
	Host struct{ Echo *echo.Echo }
)

func Factory() *echo.Echo {
	serve := echo.New()
	serve.Static("", "./public/")
	serve.Use(lang.Initialize())
	serve.Use(controller.Initialize())
	serve.Pre(middleware.RemoveTrailingSlash())
	return serve
}

func Run() {
	serve := Factory()
	// hosts := map[string]*Host{}
	storage.Connect(storage.Default())
	app.Run(serve)
	admin.Run(serve)

	// hosts[globals.Env.Host + globals.Env.Port] = &Host{ app.Run(Factory()) }
	// hosts["admin." + globals.Env.Host + globals.Env.Port] = &Host{ admin.Run(Factory()) }

	// serve.Any("/*", func(ctx echo.Context) (err error) {
	// 	req := ctx.Request()
	// 	res := ctx.Response()
	// 	host := hosts[req.Host]

	// 	if host == nil {
	// 		err = echo.ErrNotFound
	// 	} else {
	// 		host.Echo.ServeHTTP(res, req)

	// 		if globals.Env.GOENV == "development" {
	// 			built := hosts["localhost" + globals.Env.Port].Echo.Routes()
	// 			data, _ := json.MarshalIndent(built, "", "    ")
	// 			os.WriteFile("./build/routes.json", data, 0644)
	// 		}
	// 	}
	// 	return err
	// })

	serve.Logger.Fatal(serve.Start(globals.Env.Port))
}
