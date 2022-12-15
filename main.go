package main

import (
	"fmt"
	"go-redis-prac/redishelper"
	"time"
)

func main() {
	redishelper.New()

	err := redishelper.Set("test_key", "test_value", 1*time.Hour)
	if err != nil {
		fmt.Println(err)
	}

	val, err := redishelper.Get("test_key")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("test_key value : %v", val)
}
