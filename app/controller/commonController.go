package controller

import "blog/app/service"

func NewUserController() *UserController {
	user := new(UserController)
	user.UserService = service.NewUserService()
	return user
}
