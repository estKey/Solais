package route

import (
	"Solais/utils/log"
	"Solais/utils/rdb"

	"fmt"
	fhttp "github.com/valyala/fasthttp"
)

func Auth(ctx *fhttp.RequestCtx){
	log.Info.Println("Incoming login via Auth2.0")
	rdb.Set("", "user", "estKey")
	log.Info.Println("Hello", rdb.Get("", "user"))
	fmt.Fprintf(ctx, "Hello %v", rdb.Get("", "user"))
}
