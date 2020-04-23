package engine

import (
	"github.com/astaxie/beego/config"
)

// global config
var configer config.Configer

//starter
func Start()  {
	// loading config
	Configer, e := loadConfig()
	if e != nil {
		panic("loading failed...")
	}
	configer = Configer



}