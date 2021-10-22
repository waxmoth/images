package routers

import (
	"github.com/gin-gonic/gin"
	"image-functions/src/api/images"
)

func Routers() *gin.Engine {
	gin.DisableConsoleColor()

	r := gin.Default()
	r.GET("/image", images.GetImage)

	return r
}
