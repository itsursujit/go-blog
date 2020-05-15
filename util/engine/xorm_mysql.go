package engine

import (
	"blog/config"
	"blog/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/core"
	"xorm.io/xorm"
)
//实例化引擎
func NewEngine()*xorm.Engine{
	//获取配置
	_c := config.NewConfig(nil).GetMapString("mysql")
	dbUrl := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", _c["user"], _c["password"], _c["host"], _c["port"], _c["db_name"], _c["db_params"])
	//初始化引擎
	engine, err := xorm.NewEngine("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	//引擎配置
	engine.ShowSQL(true)
	engine.SetMapper(core.GonicMapper{})//映射规则
	//同步结构
	errSync := engine.Sync2(
		new(model.User),
		)
	if errSync != nil {
		panic(errSync)
	}
	return engine
}