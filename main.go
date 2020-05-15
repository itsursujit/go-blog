package main

import (
	"blog/route"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main(){
	//初始化iris对象
	app := iris.New()
	//实例化一个viper对象
	// cnf := config.NewConfig(nil).GetMapString("server")
	//设置日志级别
	// app.Logger().SetLevel(cnf["log_level"])
	//设置任何http请求恐慌恢复
	app.Use(recover.New())
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	//设置打印日志到终端
	// app.Use(logger.New())
	//注册路由
	route.NewRoute(app)
	//运行该对象
	app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithConfiguration(iris.Configuration{
			DisableStartupLog:                 false,
			DisableInterruptHandler:           false,
			DisablePathCorrection:             false,
			EnablePathEscape:                  false,
			FireMethodNotAllowed:              true, //请求方式错误会响应错误
			DisableBodyConsumptionOnUnmarshal: false,
			DisableAutoFireStatusCode:         false,
			TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
			Charset:                           "UTF-8",
		}),
	)
}
