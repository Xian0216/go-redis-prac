package main

import (
	"go-redis-prac/redishelper"
	"go-redis-prac/src/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	redishelper.New()

	app := fiber.New(fiber.Config{})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Use(recover.New())

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics Page"}))
	router.Set(app)
	log.Fatal(app.Listen(":3000"))
}
