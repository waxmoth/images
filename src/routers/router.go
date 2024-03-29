package routers

import (
	"github.com/gin-gonic/gin"
	"image-functions/src/api/images"
	"image-functions/src/middlewares"
	"os"
)

// Routers the api routers
func Routers() *gin.Engine {
	gin.DisableConsoleColor()

	r := gin.New()
	if os.Getenv("GIN_MODE") != "test" {
		r.Use(gin.Logger())
	}

	r.Use(
		gin.Recovery(),
		middlewares.AuthMiddleware(),
		middlewares.StorageMiddleware(),
	)

	r.GET("/image", images.GetImage)
	r.POST("/image", images.UploadImage)

	return r
}
