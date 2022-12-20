package controller

import "github.com/gofiber/fiber/v2"

type IRedisController interface {
	Get(c *fiber.Ctx) error
	Add(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
