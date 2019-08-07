package main

import (
	"Solais/api"
	"Solais/route"
	"Solais/utils/log"

	frouter "github.com/buaazp/fasthttprouter"
	fhttp "github.com/valyala/fasthttp"
)

type Config struct {
	id   int
	port string
}

func main() {
	var Server = Config{
		id:   1,
		port: ":8080",
	}
	router := frouter.New()
	router.GET("/", api.Route(route.Index))
	router.GET("/auth", api.Route(route.Auth))
	router.GET("/node", api.Route(route.Node))
	router.GET("/room", api.Route(route.Room))
	log.Info.Println("Start Server on Port:", Server.port)
	log.Error.Fatal(fhttp.ListenAndServe(Server.port, router.Handler))
}
