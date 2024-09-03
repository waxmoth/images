package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "image-functions/doc/api"
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

	r.Use(gin.Recovery())
	api := r.Group("/api")
	{
		api.Use(middlewares.AuthMiddleware())
		imageApi := api.Group("/image")
		{
			imageApi.Use(middlewares.StorageMiddleware())
			imageApi.GET("", images.GetImage)
			imageApi.POST("", images.UploadImage)
		}
	}

	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
