package main

import (
	"github.com/gin-gonic/gin"

	"qr-api/qr-generator/routes"
)

func main() {
	router := gin.Default()
	route.RoutesConfig(router)
	router.Run(":8080")
}
