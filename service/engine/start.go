package engine

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"net/http"
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

	// loading tcp service
	e = loadTcpService()
	if e != nil {
		panic("loading failed...")
	}
}


// loading config
func loadConfig() (Configer config.Configer,e error) {
	configer, e := config.NewConfig("ini", "./config/config.ini")
	if e != nil {
		fmt.Println("service load config fail....err:",e)
		return nil, e
	}

	return configer,nil
}

// loading tcp service
func loadTcpService() (e error) {
	fmt.Println("starting service...")

	addr := configer.String("service_host") + ":" + configer.String("service_port")
	router()
	fmt.Println("service starting ok...")

	err := http.ListenAndServe(addr, nil)

	if e != nil {
		fmt.Printf("service load tcp listen %s fail....err=%s",addr,e)
		return err
	}
	return nil
}

