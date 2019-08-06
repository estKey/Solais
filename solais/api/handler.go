package api

import (
	"log"
	"time"
	fhttp "github.com/valyala/fasthttp"
	
)

func Channel(handler fhttp.RequestHandler) fhttp.RequestHandler {
	return fhttp.RequestHandler(
		func (ctx *fhttp.RequestCtx) {
			start:= time.Now()
			handler(ctx)
			log.Println("New Connection from ", ctx.RemoteAddr(), "path ", (string)(ctx.RequestURI()), "timeuse(ns) ", time.Since(start).Nanoseconds()/1000000)
		})
}
