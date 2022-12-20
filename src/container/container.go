package container

import (
	"go-redis-prac/redishelper"
	"go-redis-prac/src/controller"
	"go-redis-prac/src/router"
)

type Injectors struct {
	Router router.IRouter
}

var _injectors = &Injectors{}

func Register() {
	_injectors.Router = router.NewRouter(controller.NewController(
		controller.NewRedisController(redishelper.GetClient()),
	))
}

func GetInjector() *Injectors {
	return _injectors
}
