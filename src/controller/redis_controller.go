package controller

import "github.com/gofiber/fiber/v2"

type IRedisController interface {
	Get(c *fiber.Ctx) error
	Add(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type redisController struct {
}

func NewRedisController() IRedisController {
	return &redisController{}
}

// Add implements IRedisController
func (*redisController) Add(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Delete implements IRedisController
func (*redisController) Delete(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Edit implements IRedisController
func (*redisController) Edit(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Get implements IRedisController
func (*redisController) Get(c *fiber.Ctx) error {
	panic("unimplemented")
}
