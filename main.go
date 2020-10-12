package main

import (
	"github.com/gin-gonic/gin"
	conf "github.com/micheleriva/gauguin/config"
	controller "github.com/micheleriva/gauguin/controller"
)

var router = gin.Default()

func init() {
	configuration := conf.ReadConfigFile()
	routes := configuration.Routes

	for _, route := range routes {
		router.GET(route.Path, func(c *gin.Context) {
			controller.HandleRoute(c, route)
		})
	}

	router.Static("/public", "./public")
}

func main() {
	router.Run()
}
