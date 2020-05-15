package controller

import (
	"blog/app/config"
	"blog/app/model"
	"blog/app/service"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	UserService *service.UserService
}

func (u *UserController) Login(ctx iris.Context) {
	jwtInfo := u.GetJWT(ctx)
	ctx.JSON(jwtInfo)
}

func (u *UserController) ChangePwd(ctx iris.Context) {

}

func (u *UserController) Index(ctx iris.Context) {
	//获取token内的信息
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token)
	ctx.JSON(jwtInfo)
}

func (u *UserController) Add(ctx iris.Context) {
	data := model.User{
		Name:     "任少鹏",
		Avatar:   "http://renshaopeng.com/img/index/me1.jpg",
		Gender:   1,
		Phone:    "13363731750",
		Password: "E10ADC3949BA59ABBE56E057F20F883E",
		Email:    "abc@163.com",
	}
	u.UserService.AddUser(data)
	ctx.JSON("Created")
}

func (u *UserController) Detail(ctx iris.Context) {

}

func (u *UserController) Update(ctx iris.Context) {

}

func (u *UserController) Delete(ctx iris.Context) {

}

func (u *UserController) GetJWT(ctx iris.Context) string {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	key := config.NewConfig(nil).GetValue("jwt_signature_key")
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
	//ctx.JSON(tokenString)
}
