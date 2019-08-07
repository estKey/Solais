package route

import (
	"fmt"
	fhttp "github.com/valyala/fasthttp"
)

func Node(ctx *fhttp.RequestCtx) {
	fmt.Fprint(ctx, "Node\n")
}