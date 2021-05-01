package logutils

import (
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func DebugMiddleWare(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		logrus.Debugf("%s %s %s", ctx.Method(), ctx.RequestURI(), ctx.PostBody())

		next(ctx)
	}
}
