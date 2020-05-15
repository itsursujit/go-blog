package service

import "blog/util/engine"

func NewUserService()*UserService{
	user := new(UserService)
	user.Engine = engine.NewEngine()
	return user
}
