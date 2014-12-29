package cache

import (
	"github.com/astaxie/beego/cache"
)

var GlobalCache cache.Cache

func init() {
	GlobalCache, _ = cache.NewCache("memory", `{"interval":3600}`)
}
