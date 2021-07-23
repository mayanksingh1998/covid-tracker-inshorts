package cache

import (
	//"encoding/json"
	//"fmt"
	//goRedis "github.com/go-redis/redis/v8"
	//"strings"
	"github.com/go-redis/redis/v8"

	"time"
)

const (
	setCacheContextTimeout  = 10 * time.Millisecond
	getCacheContextTimeout  = 10 * time.Millisecond
	mGetCacheContextTimeout = 100 * time.Millisecond
	delCacheContextTimeout  = 10 * time.Millisecond
	CacheKeyArgPrefix       = "#"
	KeySliceLimitSize       = 51
	Hour                    = time.Hour
	Day                     = 24 * Hour
	Week                    = 7 * Day
	Month                   = 30 * Day
)

type ICache interface {

}

type ICacheClient interface {

}

type Cache struct {
	client ICacheClient
}

func CacheInitializer() *redis.Client {
	var rdb *redis.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
