package service

import (
	"blog/model"
	"fmt"
	"xorm.io/xorm"
)

type UserService struct {
	Engine *xorm.Engine
}

func(u *UserService)AddUser(data model.User){
	//插入数据
	res,err := u.Engine.Insert(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
