package container

import (
	"go-redis-prac/redishelper"
	"go-redis-prac/src/controller"
	"go-redis-prac/src/router"
)

type Injectors struct {
	Ctrl   controller.IController
	Router router.IRouter
}

var _injectors = &Injectors{}

func Register() {
	_injectors.Ctrl = controller.NewController(
		controller.NewRedisController(redishelper.GetClient()),
	)
	_injectors.Router = router.NewRouter(_injectors.Ctrl)
}

func GetInjector() *Injectors {
	return _injectors
}
