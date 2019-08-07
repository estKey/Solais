package api

import (
	"Solais/utils/log"
	"time"
	fhttp "github.com/valyala/fasthttp"
	
)

func Route(handler fhttp.RequestHandler) fhttp.RequestHandler {
	return fhttp.RequestHandler(
		func (ctx *fhttp.RequestCtx) {
			start:= time.Now()
			handler(ctx)
			log.Info.Println("New Connection from", ctx.RemoteAddr(), "path:", (string)(ctx.RequestURI()), "timeuse(ns):", time.Since(start).Nanoseconds())
		})
}
