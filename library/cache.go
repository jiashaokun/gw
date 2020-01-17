package library

import (
	"fmt"
	"time"

	"gw/conf"

	"github.com/go-redis/redis/v7"
)

var rds *redis.Client

func init() {
	addr := fmt.Sprintf("%s:%s", conf.Cache["host"], conf.Cache["port"])
	rds = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "",
		DB:           0,
		PoolSize:     conf.CachePoll,
		MinIdleConns: conf.CacheMinIdleConns,
	})
}

//set
func SetCache(key, val string, tm int) error {
	if tm == 0 {
		return nil
	}

	timeOut := time.Duration(tm) * time.Second
	if err := rds.Set(key, val, timeOut).Err(); err != nil {
		return err
	}

	return nil
}

//get
func GetCache(key string) string {
	val, err := rds.Get(key).Result()
	if err != nil || err == redis.Nil {
		return ""
	}

	return val
}

//Incr
func Incr(key string) {
	rds.Incr(key)
	//过期的key保存一天时间
	expTime := 86400 * time.Second
	rds.Expire(key, expTime)
}

//HSet
func HSet(key string, f string, v interface{}) {
	rds.HSet(key, f, v)
}

//HGetNx
func HGet(key, f string) string {
	c := rds.HGet(key, f)
	res, err := c.Result()
	if err != nil {
		return ""
	}
	return res
}
