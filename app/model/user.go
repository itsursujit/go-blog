package model

type User struct {
	Id         int64
	Name       string `xorm:"varchar(255) notnull"`
	Avatar     string `xorm:"varchar(255)"`
	Gender     int8
	Phone      string `xorm:"char(11) notnull unique 'phone'"`
	Password   string `xorm:"char(32) notnull"`
	Email      string `xorm:"varchar(255) notnull unique 'email'"`
	CreateTime int64  `xorm:"created"`
	UpdateTime int64  `xorm:"updated"`
	DeleteTime int64  `xorm:"deleted"`
}
