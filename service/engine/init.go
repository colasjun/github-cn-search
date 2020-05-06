package engine

import (
	cache_store "github-cn-search/service/cache-store"
	"github.com/astaxie/beego/config"
)

func loadInit()  (e error){
	configer, e := config.NewConfig("ini", "./config/init.ini")
	if e != nil {
		return e
	}

	diy, e := configer.DIY("cache")
	if e != nil {
		return e
	}

	for k,v := range diy.(map[string]string) {
		cache_store.CACHE.Set(k,v,0)
	}

	return nil
}
