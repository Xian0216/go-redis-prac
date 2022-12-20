package router

import (
	"go-redis-prac/src/controller"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

type IRouter interface {
	New()
}

type router struct {
	ctrl controller.IController
}

func NewRouter(ctrl controller.IController) IRouter {
	return &router{
		ctrl: ctrl,
	}
}

func (r router) New() {
	app := fiber.New(fiber.Config{})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(recover.New())
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics Page"}))

	redisRouter := app.Group("/redis")
	{
		redisRouter.Get("/:key", r.ctrl.Redis().Get)
		redisRouter.Post("/:key/:value", r.ctrl.Redis().Set)
		redisRouter.Delete("/:key/:value", r.ctrl.Redis().Del)
	}

	log.Fatal(app.Listen(":3000"))
}
