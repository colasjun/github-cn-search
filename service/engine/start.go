package engine

import (
	cache_store "github-cn-search/service/cache-store"
	"github-cn-search/service/common"
	"github-cn-search/service/config"
	"github-cn-search/service/tcp"
)

//starter
func Start()  {
	// loading config
	e := config.LoadConfig()
	if e != nil {
		panic("loading config " + common.FailMsg.PanicMsg)
	}

	// loading cache server
	e = cache_store.LoadCache()
	if e != nil {
		panic("loading cache " + common.FailMsg.PanicMsg)
	}

	// loading tcp service
	e = tcp.LoadTcpService()
	if e != nil {
		panic("loading service " + common.FailMsg.PanicMsg)
	}
}

