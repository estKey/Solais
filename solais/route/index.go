package route

import (
	"fmt"
	"Solais/utils/log"
	fhttp "github.com/valyala/fasthttp"
)

func Index(ctx *fhttp.RequestCtx){
	fmt.Fprint(ctx, "Welcome to Solais!\n")
	log.Info.Println("Now Visiting Index Page")
	// reflect.TypeOf(d)
}
