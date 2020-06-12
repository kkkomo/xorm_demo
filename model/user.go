package model

import (
	"github.com/micro/go-micro/util/log"
	"time"
)

type (
	User struct {
		Id int64 `xorm:"id"`
		Name string `xorm:"name"`
		Age int64 `xorm:"age"`
		Password string `xorm:"password"`
		Birthday time.Time `xorm:"birthday"`
	}
)

func (User) TableName() string {
	return "users"
}

func GetUserByName(name string) User {
	user := User{}
	has, err := engine.Table("users").Where("name=?", name).Get(&user)
	if err != nil {
		log.Error("select user error:", err)
	}
	log.Info("res :", has)
	return user
}