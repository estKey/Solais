package api

import (
	"Solais/utils/log"
	"time"
	"github.com/valyala/fasthttp"
)

func Route(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(
		func (ctx *fasthttp.RequestCtx) {
			start:= time.Now()
			handler(ctx)
			log.Info.Println("New Connection from", ctx.RemoteAddr(), "path:", (string)(ctx.RequestURI()), "timeuse(ns):", time.Since(start).Nanoseconds())
		})
}

// func Request(){
// 	req := fasthttp.AcquireRequest()
// 	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	
// 	// 默认是application/x-www-form-urlencoded
// 	req.Header.SetContentType("application/json")
// 	req.Header.SetMethod("POST")
	
// 	req.SetRequestURI(url)
	
// 	requestBody := []byte(`{"request":"test"}`)
// 	req.SetBody(requestBody)
// }

// func Response(url string) {
// 	url := `http://httpbin.org/post?key=123`
// 	resp := fasthttp.AcquireResponse()
// 	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

// 	if err := fasthttp.Do(req, resp); err != nil {
// 			fmt.Println("请求失败:", err.Error())
// 			return
// 	}

// 	b := resp.Body()

// 	fmt.Println("result:\r\n", string(b))
// }