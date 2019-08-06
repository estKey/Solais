package route

import (
	"fmt"
	fhttp "github.com/valyala/fasthttp"
)

func Index(ctx *fhttp.RequestCtx){
	fmt.Fprint(ctx, "Welcome FastHttp!\n")
}
