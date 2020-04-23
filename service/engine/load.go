package engine

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

// loading config
func loadConfig() (Configer config.Configer,e error) {
	configer, e := config.NewConfig("ini", "./config/config.ini")
	if e != nil {
		fmt.Println("service load config fail....err:",e)
		return nil, nil
	}

	return configer,nil
}
