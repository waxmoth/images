package routers

import (
	"github.com/gin-gonic/gin"
	"image-functions/src/api/images"
	"image-functions/src/middlewares"
)

func Routers() *gin.Engine {
	gin.DisableConsoleColor()

	r := gin.Default()
	r.GET("/image", images.GetImage)
	r.Use(middlewares.JSONMiddleware())

	return r
}
