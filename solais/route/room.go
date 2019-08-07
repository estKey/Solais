package route

import (
	"fmt"
	fhttp "github.com/valyala/fasthttp"
)

func Room(ctx *fhttp.RequestCtx) {
	fmt.Fprint(ctx, "Room\n")
}