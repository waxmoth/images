package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"image-functions/src/consts"
	"net/url"
	"strings"
)

// ParseURL parse url and validate the host from the context
func ParseURL(rawURL string, ct *gin.Context) (*url.URL, error) {
	data, exist := ct.Get(consts.AuthorizedData)
	if !exist {
		return nil, errors.New("missing authorized data")
	}
	jsonData := data.(map[string]interface{})
	parsedURL, err := url.Parse(rawURL)
	if err == nil && jsonData["hosts"] != nil && strings.Contains(jsonData["hosts"].(string), parsedURL.Host) {
		return parsedURL, nil
	}
	return parsedURL, errors.New("the url is invalid")
}
