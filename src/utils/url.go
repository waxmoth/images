package utils

import (
	"github.com/gin-gonic/gin"
	"net/url"
)

// ParseURL parse url and validate the host from the context
func ParseURL(rawURL string, ct *gin.Context) (*url.URL, error) {
	// TODO verify the url host from the JWT token
	return url.Parse(rawURL)
}
