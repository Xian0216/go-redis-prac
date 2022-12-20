package router

import "github.com/gofiber/fiber/v2"

func Set(r fiber.Router) {
	redisRouter := r.Group("redis/")
	{
		redisRouter.Get(":key")
		redisRouter.Post(":key/:value")
		redisRouter.Put(":key/:value")
		redisRouter.Delete(":key/:value")
	}
}
