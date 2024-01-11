package main

import (
	"github.com/gin-gonic/gin"

	route "qr-api/qr-generator/routes"

	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}

	router.Use(cors.New(config))

	route.RoutesConfig(router)
	router.Run(":8080")
}
