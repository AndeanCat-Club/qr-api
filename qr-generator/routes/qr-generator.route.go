package route

import (
	"qr-api/qr-generator/controllers"
	"github.com/gin-gonic/gin"

)

func RoutesConfig(router *gin.Engine) {
	router.POST("/post-generate-qr", controller.GenerateQR)
}