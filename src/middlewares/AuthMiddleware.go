package middlewares

import (
	"github.com/gin-gonic/gin"
	"image-functions/src/consts"
	"image-functions/src/services/auth"
	"net/http"
	"os"
	"strings"
)

// AuthMiddleware Auth the HTTP header, return 401 error if unauthorized
func AuthMiddleware() gin.HandlerFunc {
	return func(ct *gin.Context) {
		reqToken := ct.GetHeader("Authorization")
		if reqToken == "" || !strings.Contains(reqToken, "Bearer ") {
			ct.Data(http.StatusUnauthorized, "application/json", []byte("Unauthorized"))
			ct.Abort()
			return
		}
		tokenString := reqToken[len("Bearer "):]
		var authService auth.Auth = &auth.JWTService{
			Key: os.Getenv("AUTH_KEY"),
		}
		token, err := authService.Decode(tokenString)
		if err != nil {
			ct.Data(http.StatusUnauthorized, "application/json", []byte("Unauthorized"))
			ct.Abort()
			return
		}
		ct.Set(consts.AuthorizedData, token)
		ct.Next()
	}
}
