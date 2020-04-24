package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

var CONF *CONFOBJ

type CONFOBJ struct {
	C config.Configer
}

func (c *CONFOBJ) Get(key string) (s string) {
	return c.C.String(key)
}

// loading cache
func LoadConfig() (e error) {
	configer, e := config.NewConfig("ini", "./config/config.ini")
	if e != nil {
		fmt.Println("service load config fail....err:",e)
		return e
	}

	CONF = &CONFOBJ{
		C:configer,
	}

	return nil
}
