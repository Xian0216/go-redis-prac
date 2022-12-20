package controller

type IController interface {
	Redis() IRedisController
}

type controller struct {
	redisCtrl IRedisController
}

func NewController(redisCtrl IRedisController) IController {
	return &controller{
		redisCtrl: redisCtrl,
	}
}

func (c *controller) Redis() IRedisController {
	return c.redisCtrl
}
