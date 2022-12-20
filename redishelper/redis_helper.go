package redishelper

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var _ctx = context.Background()
var rdb *redis.Client

func New() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Set(key string, value interface{}, duration time.Duration) error {
	return rdb.Set(_ctx, key, value, duration).Err()
}

func Get(key string) (string, error) {
	return rdb.Get(_ctx, key).Result()
}
