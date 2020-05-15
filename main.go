package main

import (
	"blog/app/route"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Handle("GET", "/", func(ctx iris.Context) {
		_, _ = ctx.HTML("<h1>Welcome</h1>")
	})

	route.NewRoute(app)
	_ = app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithConfiguration(iris.Configuration{
			DisableStartupLog:                 false,
			DisableInterruptHandler:           false,
			DisablePathCorrection:             false,
			EnablePathEscape:                  false,
			FireMethodNotAllowed:              true,
			DisableBodyConsumptionOnUnmarshal: false,
			DisableAutoFireStatusCode:         false,
			TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
			Charset:                           "UTF-8",
		}),
	)
}
