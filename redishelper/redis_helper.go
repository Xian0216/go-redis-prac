package redishelper

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var _ctx = context.Background()
var _rdb *redis.Client

func New() {
	_rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetClient() *redis.Client {
	return _rdb
}

func Set(rdb *redis.Client, key string, value interface{}, expiration time.Duration) error {
	return rdb.Set(_ctx, key, value, expiration).Err()
}

func Get(rdb *redis.Client, key string) (string, error) {
	return rdb.Get(_ctx, key).Result()
}

func Del(rdb *redis.Client, key ...string) error {
	return rdb.Del(_ctx, key...).Err()
}
