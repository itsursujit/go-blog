package controller

import "blog/service"

func NewUserController()*UserController{
	user := new(UserController)
	user.UserService = service.NewUserService()
	return user
}
