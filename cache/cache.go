package cache

import (
	"context"
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
	getCachePrefix() string
	getCacheContext() context.Context
	GetKey(k string, kwargs *map[string]string) string
	Get(k string) (string, error)
	Set(k string, i interface{}, ttl time.Duration) error
	Del(keys ...string) error
	HSet(k string, data map[string]interface{}, ttl time.Duration) error
	HMGet(k string, keys []string) ([]interface{}, error)
	MSet(data map[string]interface{}, ttl time.Duration) error
	MGet(keys []string) ([]interface{}, error)
	HDel(key string, fields []string) error
}

type ICacheClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	MGet(ctx context.Context, keys []string) ([]interface{}, error)
	HSet(ctx context.Context, key string, values map[string]interface{}) error
	HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error)
	MSet(ctx context.Context, values map[string]interface{}) error
	Del(ctx context.Context, keys ...string) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
	HDel(ctx context.Context, key string, fields ...string) error
}

type Cache struct {
	client ICacheClient
}

func CacheInitializer() *redis.Client {
	//cacheMode := configInstance.GetServiceCacheMode()
	//if cacheMode == config.NodeCacheMode {
	//	cache.client = redis
	//}
	//return cache
	var rdb *redis.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
//func set(c *redis.Client, key string, value interface{}) error {
//	ac := context.Background()
//	p, err := json.Marshal(value)
//	if err != nil {
//		return err
//	}
//	return c.Set(key, p)
//	//var a = c.Set(ac, key, p, time.Second)
//	//return c.Set(ac, key, p, time.Second)
//}
//
//func get(c *redis.Client, key string, dest interface{}) error {
//	p := c.Get(context.Context, key)
//	if err != nil {
//		return err
//	}
//	return json.Unmarshal(p, dest)
//}
//func (c *Cache) getCachePrefix() string {
//	return c.config.GetServiceCachePrefix()
//}
//
//func (c *Cache) getCacheContext() context.Context {
//	return context.Background()
//}
//
//func MarshalBinary(m interface{}) ([]byte, error) {
//	return json.Marshal(m)
//}
//
//func UnmarshalBinary(data []byte, m interface{}) error {
//	if err := json.Unmarshal(data, m); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (c *Cache) GetKey(k string, kwargs *map[string]string) string {
//	if kwargs != nil {
//		for ka, v := range *kwargs {
//			k = strings.ReplaceAll(k, CacheKeyArgPrefix+ka, v)
//		}
//	}
//
//	return fmt.Sprintf("%s_%s", c.getCachePrefix(), k)
//}
//
//func (c *Cache) Get(k string) (string, error) {
//	ctx, cancel := context.WithTimeout(c.getCacheContext(), getCacheContextTimeout)
//	defer cancel()
//
//	v, err := c.client.Get(ctx, k)
//	if err != nil {
//		c.logger.Error(fmt.Sprintf("Cache key: %s get error: %s ", k, err.Error()))
//	}
//
//	return v, err
//}
//
//func (c *Cache) Set(k string, i interface{}, expireDuration time.Duration) error {
//	ctx, cancel := context.WithTimeout(c.getCacheContext(), setCacheContextTimeout)
//	defer cancel()
//
//	marshalBinary, err := MarshalBinary(i)
//	if err != nil {
//		c.logger.Error(fmt.Sprintf("Cache key: %s marshal error: %s ", k, err.Error()))
//		return err
//	}
//
//	err = c.client.Set(ctx, k, marshalBinary, expireDuration)
//	if err != nil && err != goRedis.Nil {
//		c.logger.Error(fmt.Sprintf("Cache key: %s set error: %s ", k, err.Error()))
//	}
//
//	return err
//}
//
//func (c *Cache) Del(keys ...string) error {
//	ctx, cancel := context.WithTimeout(c.getCacheContext(), delCacheContextTimeout)
//	defer cancel()
//
//	err := c.client.Del(ctx, keys...)
//	if err != nil {
//		c.logger.Error(fmt.Sprintf("Cache keys: %s delete error: %s ", keys, err.Error()))
//	}
//
//	return err
//}
//
//func (c *Cache) HSet(k string, data map[string]interface{}, ttl time.Duration) error {
//	ctx, cancel := context.WithTimeout(c.getCacheContext(), setCacheContextTimeout)
//	defer cancel()
//
//	marshalledData := make(map[string]interface{})
//	for k, v := range data {
//		d, err := MarshalBinary(v)
//		if err != nil {
//			return err
//		}
//		marshalledData[k] = d
//	}
//
//	err := c.client.HSet(ctx, k, marshalledData)
//	if err != nil && err != goRedis.Nil {
//		c.logger.Error(fmt.Sprintf("Cache key: %s hset error: %s", k, err.Error()))
//		return err
//	}
//
//	err = c.client.Expire(ctx, k, ttl)
//	if err != nil {
//		c.logger.Error(fmt.Sprintf("Cache key: %s expire error: %s", k, err.Error()))
//	}
//
//	return err
//}
//
//func (c *Cache) HMGet(k string, keys []string) ([]interface{}, error) {
//	ctx, cancel := context.WithTimeout(c.getCacheContext(), getCacheContextTimeout)
//	defer cancel()
//
//	v, err := c.client.HMGet(ctx, k, keys...)
//	if err != nil {
//		c.logger.Error(fmt.Sprintf("Cache key: %s keys: %s hmget error: %s  ", k, keys, err.Error()))
//	}
//
//	return v, err
//}
//
//func (c *Cache) MSet(data map[string]interface{}, ttl time.Duration) error {
//	ctx, cancel := context.WithTimeout(c.getCacheContext(), setCacheContextTimeout)
//	defer cancel()
//
//	marshalledData := make(map[string]interface{})
//	for k, v := range data {
//		d, err := MarshalBinary(v)
//		if err != nil {
//			c.logger.Error(fmt.Sprintf("Cache key: %s marshal error: %s ", k, err.Error()))
//			return err
//		}
//		marshalledData[k] = d
//	}
//
//	err := c.client.MSet(ctx, marshalledData)
//	if err != nil {
//		c.logger.Error(fmt.Sprintf("Cache mset error: %s ", err.Error()))
//	} else {
//		for key := range marshalledData {
//			err = c.client.Expire(ctx, key, ttl)
//			if err != nil {
//				c.logger.Error(fmt.Sprintf("Cache key: %s expire error: %s", key, err.Error()))
//			}
//		}
//	}
//
//	return err
//}
//
//func (c *Cache) MGet(keys []string) ([]interface{}, error) {
//	ctx, cancel := context.WithTimeout(c.getCacheContext(), mGetCacheContextTimeout)
//	defer cancel()
//
//	v, err := c.client.MGet(ctx, keys)
//	if err != nil {
//		c.logger.Error(fmt.Sprintf("Cache keys: %s mget error: %s ", keys, err.Error()))
//	}
//
//	return v, err
//}
//
//func (c *Cache) HDel(key string, fields []string) error {
//	ctx, cancel := context.WithTimeout(c.getCacheContext(), delCacheContextTimeout)
//	defer cancel()
//
//	err := c.client.HDel(ctx, key, fields...)
//	if err != nil {
//		c.logger.Error(fmt.Sprintf("Cache key: %s hdel error: %s fields: %s ", key, err.Error(), fields))
//	}
//
//	return err
//}
