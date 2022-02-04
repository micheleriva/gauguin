package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	config "github.com/micheleriva/gauguin/config"
	controller "github.com/micheleriva/gauguin/controller"
)

var router = gin.Default()

func init() {
	if config.ConfigError == nil {
		router.Static("/public", "./public")
		router.StaticFile("/favicon.ico", "./assets/favicon.ico")
		router.NoRoute(controller.HandleRoutes)
	} else {
		cwd, _ := os.Getwd()
		log.Printf("error occurred while trying to read Gauguin configuration: %v\n", config.ConfigError)
		log.Printf("we've tried to read the following configuration file: %s\n", fmt.Sprintf("%s/gauguin.yaml", cwd))

		router.NoRoute(controller.ConfigError)
	}
}

func main() {
	router.Run()
}
