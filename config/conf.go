package config

import (
	"github.com/BurntSushi/toml"
	"github.com/micro/go-micro/util/log"
)

var Config = &conf{}

func init() {
	file := "config/dev.toml"
	if _, err := toml.DecodeFile(file, Config); err != nil {
		panic(err)
	}
	log.Info("Mysql config : ", Config.Mysql)
}

type (
	conf struct {
		Mysql *Mysql `toml:"mysql"`
	}
	Mysql struct {
		Dns string `toml:"dns"`
		MaxIdleConn int `toml:"maxIdleConn"`
		MaxOpenConn int `toml:"maxOpenConn"`
		MaxLifetime int `toml:"maxLifetime"`
	}
)
