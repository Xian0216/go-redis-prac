package controller

import (
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func Test_redisController_Get(t *testing.T) {
	expectKey := "justKey"
	expectValue := "justValue"

	fctx := &fasthttp.RequestCtx{}
	fctx.URI().QueryArgs().Set("key", expectKey)

	ctx := givenfiberApp(fctx)

	rdb, rdmock := redismock.NewClientMock()
	rdmock.ExpectGet(expectKey).SetVal(expectValue)

	err := NewRedisController(rdb).Get(ctx)
	assert.NoError(t, err)
	assert.EqualValues(t, expectValue, string(ctx.Response().Body()))
}

func givenfiberApp(fctx *fasthttp.RequestCtx) *fiber.Ctx {
	return fiber.New().AcquireCtx(fctx)
}
