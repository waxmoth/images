package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"image-functions/src/consts"
	"net/url"
)

// ParseURL parse url and validate the host from the context
func ParseURL(rawURL string, ct *gin.Context) (*url.URL, error) {
	data, exist := ct.Get(consts.AuthorizedData)
	if !exist {
		return nil, errors.New("missing authorized data")
	}
	jsonData := data.(map[string]interface{})
	parsedURL, err := url.Parse(rawURL)
	if err == nil && jsonData["host"] != nil && jsonData["host"].(string) == parsedURL.Host {
		parsedURL.Host = jsonData["host"].(string)
		return parsedURL, nil
	}
	return nil, errors.New("the url is invalid")
}
