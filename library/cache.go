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
func Set(key, val string, tm int) error {
	timeOut := time.Duration(tm) * time.Second
	if err := rds.Set(key, val, timeOut).Err(); err != nil {
		return err
	}

	return nil
}

//get
func Get(key string) (string, error) {
	val, err := rds.Get(key).Result()
	if err != nil || err == redis.Nil {
		return "", nil
	}

	return val, nil
}
