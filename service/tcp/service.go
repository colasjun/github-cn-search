package tcp


import (
	"github-cn-search/service/config"
	"net/http"
	"fmt"
	)

// loading tcp service
func LoadTcpService() (e error) {
	fmt.Println("starting service...")

	addr := config.CONF.Get("SERVICE_HOST") + ":" + config.CONF.Get("SERVICE_PORT")

	// router...
	router()

	fmt.Println("service starting ok...")

	err := http.ListenAndServe(addr, nil)

	if e != nil {
		fmt.Printf("service load tcp listen %s fail....err=%s",addr,e)
		return err
	}
	return nil
}
