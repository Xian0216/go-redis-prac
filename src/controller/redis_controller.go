package controller

import (
	"go-redis-prac/redishelper"
	"time"

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

func (r *redisController) Get(c *fiber.Ctx) error {
	k := c.Params("key")
	if k == "" {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}
	v, err := redishelper.Get(r.rdb, k)
	if err != nil {
		return c.SendStatus(fiber.ErrInternalServerError.Code)
	}

	return c.Status(fiber.StatusOK).SendString(v)
}

func (r *redisController) Set(c *fiber.Ctx) error {
	k, v := c.Params("key"), c.Params("value")
	if k == "" || v == "" {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}
	if err := redishelper.Set(r.rdb, k, v, 1*time.Hour); err != nil {
		return c.SendStatus(fiber.ErrInternalServerError.Code)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (r *redisController) Del(c *fiber.Ctx) error {
	k := c.Params("key")
	if k == "" {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}
	if err := redishelper.Del(r.rdb, k); err != nil {
		return c.SendStatus(fiber.ErrInternalServerError.Code)
	}
	return c.SendStatus(fiber.StatusOK)
}
