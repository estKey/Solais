package route

import (
	"log"
	fhttp "github.com/valyala/fasthttp"
)

func Auth(ctx *fhttp.RequestCtx){
	log.Println("Incoming login via Auth2.0")
}
