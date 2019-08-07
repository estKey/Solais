package route

import (
	"fmt"
	"log"
	fhttp "github.com/valyala/fasthttp"
)

func Index(ctx *fhttp.RequestCtx){
	fmt.Fprint(ctx, "Welcome to Solais!\n")
	log.Println("Index Page")
	// reflect.TypeOf(d)
}
