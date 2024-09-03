package api

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/gin-gonic/gin"
	"image-functions/src/utils"
)

//	@title			Image Functions
//	@version		1.0
//	@description	GoLang Image Functions

//	@contact.name	waxmoth
//	@contact.email	waxmoth@qq.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@BasePath		/api

type meta struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	ProcessAt   int64  `json:"processAt"`
	ProcessedID string `json:"processedId"`
}

type Error struct {
	Code        int    `json:"code"`
	Error       string `json:"error"`
	ProcessAt   int64  `json:"processAt"`
	ProcessedID string `json:"processedId"`
}

type SuccessResponse struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// ReturnSuccess return the success response body to client
func ReturnSuccess(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(code, SuccessResponse{
		Meta: meta{
			code,
			msg,
			utils.NowMillis(),
			createProcessedID(c),
		},
		Data: data,
	})
}

// ReturnError return error information to client
func ReturnError(code int, msg string, c *gin.Context) {
	c.JSON(code, Error{
		code,
		msg,
		utils.NowMillis(),
		createProcessedID(c),
	})
}

func createProcessedID(c *gin.Context) string {
	// Note: Get lambda request id as process id
	if lambdaContext, ok := core.GetRuntimeContextFromContext(c.Request.Context()); ok {
		return lambdaContext.AwsRequestID
	}

	u := make([]byte, 8)
	_, err := rand.Read(u)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(u)
}

// ReturnFile return the file buffer to the client
func ReturnFile(code int, contentType string, data []byte, c *gin.Context) {
	c.Data(code, contentType, data)
}
