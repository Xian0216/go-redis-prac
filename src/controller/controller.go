package controller

type IController interface {
	Redis() IRedisController
}

type controller struct {
}

func NewController() IController {
	return &controller{}
}

func (*controller) Redis() IRedisController {
	return NewRedisController()
}
