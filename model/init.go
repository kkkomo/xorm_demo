package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/micro/go-micro/util/log"
	"xorm.io/core"
	"xorm_demo/config"
)

var (
	conf = config.Config
	engine *xorm.Engine
)

func init() {
	en, err := xorm.NewEngine("mysql", conf.Mysql.Dns)
	if err != nil {
		log.Error("connect mysql error")
		panic(err)
	}
	log.Info("connect mysql success")
	en.ShowSQL(true)
	en.SetMaxIdleConns(conf.Mysql.MaxIdleConn)
	en.SetMaxOpenConns(conf.Mysql.MaxOpenConn)
	en.SetTableMapper(core.SnakeMapper{})
	/*err = en.Sync2(new(User))
	if err != nil {
		log.Error("sync table error : ", err)
		panic(err)
	}*/
	engine = en
}