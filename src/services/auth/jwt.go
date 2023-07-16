package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// JWTService the JWT service to encode and decode data
type JWTService struct {
	Key     string // The secret key to sign and verify token
	Expires int    // The expiration time, unit hours
}

// Encode data and sign it by the secret key
func (jwtService *JWTService) Encode(obj interface{}) (string, error) {
	if jwtService.Key == "" {
		return "", fmt.Errorf("key is empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"app":  "image-functions",
		"data": obj,
		"exp":  time.Now().Add(time.Duration(jwtService.Expires) * time.Hour).Unix(),
	})
	return token.SignedString([]byte(jwtService.Key))
}

// Decode data and validate it from the secret key
func (jwtService *JWTService) Decode(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtService.Key), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, err := claims.GetExpirationTime(); err != nil || exp.Unix() < time.Now().Unix() {
			return nil, fmt.Errorf("token expired")
		}
		return claims["data"], nil
	}

	return nil, err
}
