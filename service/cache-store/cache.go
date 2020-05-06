package cache_store

import (
	"fmt"
	"github-cn-search/service/config"
	"github.com/go-redis/redis"
	"time"
)

var CACHE *CACHEOBJ

type CACHEOBJ struct {
	C *redis.Client
}

func (c *CACHEOBJ) Get(key string) (s string) {
	s, e := c.C.Get(key).Result()
	if e != nil {
		fmt.Println("cache get err...err=",e)
		return ""
	}

	return s
}

func (c *CACHEOBJ) Set(key string, value string, expiration time.Duration) (s string) {
	c.C.Set(key, value, 0)
	return ""
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