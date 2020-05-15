package route

import (
	"blog/controller"
	"blog/middleware"
	"github.com/kataras/iris/v12"
)

func newBackendRoute(party iris.Party){
	handleUser(party)
}

func handleUser(party iris.Party){
	user := controller.NewUserController()
	party.Post("/login",user.Login).Name = "adminLogin"
	party.Post("/change-pwd",user.ChangePwd).Name = "changePassword"
	party.Get("/index",middleware.JWT.Serve,user.Index).Name = "userDashboard"
	party.Get("/detail",middleware.JWT.Serve,user.Detail).Name = "userDetail"
	party.Post("/add",middleware.JWT.Serve,user.Add).Name = "addUser"
	party.Delete("/delete",middleware.JWT.Serve,user.Delete).Name = "deleteUser"
	party.Put("/update",middleware.JWT.Serve,user.Update).Name = "updateUser"
}