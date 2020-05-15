package route

import (
	"github.com/kataras/iris/v12"
)

func NewRoute(app *iris.Application) {
	newBackendRoute(app.Party("/admin"))
}
