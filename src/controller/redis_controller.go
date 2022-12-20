package controller

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type IRedisController interface {
	Get(c *fiber.Ctx) error
	Set(c *fiber.Ctx) error
	Del(c *fiber.Ctx) error
}

type redisController struct {
	rdb *redis.Client
}

func NewRedisController(rdb *redis.Client) IRedisController {
	return &redisController{
		rdb: rdb,
	}
}

// Get implements IRedisController
func (*redisController) Get(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Set implements IRedisController
func (*redisController) Set(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Del implements IRedisController
func (*redisController) Del(c *fiber.Ctx) error {
	panic("unimplemented")
}
