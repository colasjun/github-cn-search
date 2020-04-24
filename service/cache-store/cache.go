package cache_store

import (
	"fmt"
	"github-cn-search/service/config"
	"github.com/go-redis/redis"
)

var CACHE *CACHEOBJ

type CACHEOBJ struct {
	C *redis.Client
}

func (c *CACHEOBJ) Get(key string) (s string) {
	//LoadCache()
	s, e := c.C.Get(key).Result()
	if e != nil {
		fmt.Println("cache err...err=",e)
		return ""
	}

	return s
}

// loading cache
func LoadCache() (e error) {
	var cacheAddr= config.CONF.Get("CACHE_HOST") + ":" + config.CONF.Get("CACHE_PORT")
	fmt.Println("Cache addr=",cacheAddr)
	var c = redis.NewClient(&redis.Options{Addr: cacheAddr,PoolSize: 1})
	//defer c.Close()

	CACHE = &CACHEOBJ{
		C:c,
	}

	return nil
}