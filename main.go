package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	config "github.com/micheleriva/gauguin/config"
	controller "github.com/micheleriva/gauguin/controller"
)

var router = gin.Default()

func init() {
	if config.ConfigError == nil {
		// router.LoadHTMLGlob("editor/*")
		router.Static("/public", "./public")
		router.StaticFile("/favicon.ico", "./assets/favicon.ico")
		// router.GET("/", editor.RenderEditor)
		router.NoRoute(controller.HandleRoutes)
	} else {
		cwd, _ := os.Getwd()
		fmt.Println("An error occurred while trying to read Gauguin configuration:")
		fmt.Println(config.ConfigError)
		fmt.Println("We've tried to read the following configuration file:")
		fmt.Println(fmt.Sprintf("%s/gauguin.yaml", cwd))

		router.NoRoute(controller.ConfigError)
	}
}

func main() {
	router.Run()
}
