package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/micheleriva/gauguin/controller"
)

var router = gin.Default()

func init() {
	router.Static("/public", "./public")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")
	router.NoRoute(controller.HandleRoutes)
}

func main() {
	router.Run()
}
