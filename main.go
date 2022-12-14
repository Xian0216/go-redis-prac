package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main(){
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	err := rdb.Set(ctx, "test_key", "test_value", 1*time.Hour).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := rdb.Get(ctx, "test_key").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("test_key value : %v", val)
}
