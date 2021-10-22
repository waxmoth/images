package api

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/gin-gonic/gin"
	"image-functions/src/utils"
	"net/http"
)

type Meta struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	ProcessAt   int64  `json:"processAt"`
	ProcessedId string `json:"processedId"`
}

type SuccessResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Meta  Meta  `json:"meta"`
	Error Error `json:"error"`
}

type Error struct {
	Name string `json:"name"`
}

func ReturnSuccess(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, SuccessResponse{
		Meta: Meta{
			code,
			msg,
			utils.NowMillis(),
			createProcessedId(c),
		},
		Data: data,
	})
}

func ReturnError(code int, msg string, c *gin.Context) {
	c.JSON(code, Meta{
		code,
		msg,
		utils.NowMillis(),
		createProcessedId(c),
	})
}

func createProcessedId(c *gin.Context) string {
	// Note: Get lambda request id as process id
	if lambdaContext, ok := core.GetRuntimeContextFromContext(c.Request.Context()); ok {
		return lambdaContext.AwsRequestID
	}

	u := make([]byte, 8)
	rand.Read(u)
	return hex.EncodeToString(u)
}

func ReturnFile(code int, contentType string, data []byte, c *gin.Context) {
	c.Data(code, contentType, data)
}
