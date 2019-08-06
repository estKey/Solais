package main

import (
	"Solais/api"
	"Solais/route"
	"log"
	frouter "github.com/buaazp/fasthttprouter"
	fhttp "github.com/valyala/fasthttp"
)

func main()  {
	var port int = 8080
	router := frouter.New()
	router.GET("/", api.Channel(route.Index))
	log.Println("Start Server on Port:", port)
	log.Fatal(fhttp.ListenAndServe(":8080", router.Handler))
}
