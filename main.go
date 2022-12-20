package main

import (
	"go-redis-prac/redishelper"
	"go-redis-prac/src/container"
)

func main() {
	redishelper.New()
	container.Register()
	container.GetInjector().Router.New()
}
