package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "qr-api/qr-generator/service"
)

func GenerateQR(c *gin.Context) {
    data := c.Query("data")

    if data == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
        return
    }

    qrPNG, err := service.GenerateQRPNG(data)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    c.Data(http.StatusOK, "image/png", qrPNG)
}