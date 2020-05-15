package controller

import (
	"blog/config"
	"blog/model"
	"blog/service"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	UserService *service.UserService
}

//登录
func(u *UserController)Login(ctx iris.Context){
	//验证登录密码
	jwtInfo := u.GetJWT(ctx)
	ctx.JSON(jwtInfo)
}

//更改密码
func(u *UserController)ChangePwd(ctx iris.Context){

}

//首页
func(u *UserController)Index(ctx iris.Context){
	//获取token内的信息
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token)
	ctx.JSON(jwtInfo)
}

//添加管理员
func(u *UserController)Add(ctx iris.Context){
	//数据验证todo...
	data := model.User{
		Name: "任少鹏",
		Avatar: "http://renshaopeng.com/img/index/me1.jpg",
		Gender: 1,
		Phone: "13363731750",
		Password: "E10ADC3949BA59ABBE56E057F20F883E",
		Email: "abc@163.com",
	}
	u.UserService.AddUser(data)
	ctx.JSON("什么情况")
}

//管理员详情
func(u *UserController)Detail(ctx iris.Context){

}

//更新管理员
func(u *UserController)Update(ctx iris.Context){

}

//删除管理员
func(u *UserController)Delete(ctx iris.Context){

}

func(u *UserController)GetJWT(ctx iris.Context)string{
	// 往jwt中写入了一对值
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	// 签名生成jwt字符串
	key := config.NewConfig(nil).GetValue("jwt_signature_key")
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
	//ctx.JSON(tokenString)
}