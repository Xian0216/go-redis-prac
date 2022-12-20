package router

import (
	"go-redis-prac/src/controller"

	"github.com/gofiber/fiber/v2"
)

func Set(r fiber.Router) {
	ctrl := controller.NewController()
	redisRouter := r.Group("redis/")
	{
		redisRouter.Get(":key", ctrl.Redis().Get)
		redisRouter.Post(":key/:value", ctrl.Redis().Set)
		redisRouter.Put(":key/:value", ctrl.Redis().Edit)
		redisRouter.Delete(":key/:value", ctrl.Redis().Delete)
	}
}
