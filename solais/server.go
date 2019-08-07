package main

import (
	"Solais/api"
	"Solais/route"

	"log"

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
	router.GET("/", api.Channel(route.Index))
	router.GET("/auth", api.Channel(route.Auth))
	router.GET("/node", api.Channel(route.Node))
	router.GET("/room", api.Channel(route.Room))
	log.Println("Start Server on Port:", Server.port)
	log.Fatal(fhttp.ListenAndServe(Server.port, router.Handler))
}
