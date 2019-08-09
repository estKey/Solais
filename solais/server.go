package main

import (
	"Solais/api"
	"Solais/route"
	"Solais/utils"
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
		port: ":9000",
	}
	router := frouter.New()
	router.GET("/", api.Route(route.Index))
	router.GET("/auth", api.Route(route.Auth))
	router.GET("/node", api.Route(route.Node))
	router.GET("/room", api.Route(route.Room))
	log.Info.Println("Starting Server on Port:", Server.port, ", Total Uptime:", utils.CalculateUptime())
	log.Error.Fatal(fhttp.ListenAndServe(Server.port, router.Handler))
	log.Info.Println("Shutting Down, Total Uptime:", utils.CalculateUptime())
}


func configureRouter(router *frouter.Router) {
}