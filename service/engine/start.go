package engine

import (
	cache_store "github-cn-search/service/cache-store"
	"github-cn-search/service/config"
	"github-cn-search/service/tcp"
)

//starter
func Start()  {
	// loading config
	e := config.LoadConfig()
	if e != nil {
		panic("loading config failed...")
	}

	// loading cache server
	e = cache_store.LoadCache()
	if e != nil {
		panic("loading config failed...")
	}

	// loading tcp service
	e = tcp.LoadTcpService()
	if e != nil {
		panic("loading service failed...")
	}
}

