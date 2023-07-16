package middlewares

import (
	"github.com/gin-gonic/gin"
	"image-functions/src/services/auth"
	"log"
	"net/http"
	"os"
)

// AuthMiddleware Auth the HTTP header, return 401 error if unauthorized
func AuthMiddleware() gin.HandlerFunc {
	return func(ct *gin.Context) {
		reqToken := ct.GetHeader("Authorization")
		if reqToken == "" {
			ct.Data(http.StatusUnauthorized, "application/json", []byte("Unauthorized"))
			ct.Abort()
			return
		}
		tokenString := reqToken[len("Bearer "):]
		var authService auth.Auth = &auth.JWTService{
			Key: os.Getenv("AUTH_KEY"),
		}
		_, err := authService.Decode(tokenString)
		if err != nil {
			log.Println(err.Error())
			ct.Data(http.StatusUnauthorized, "application/json", []byte("Unauthorized"))
			ct.Abort()
			return
		}
		ct.Next()
	}
}
