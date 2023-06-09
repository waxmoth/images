package images

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"image"
	"image-functions/src/api"
	"image-functions/src/consts"
	"image-functions/src/models/requests"
	"image-functions/src/utils"
	"image/jpeg"
	"log"
	"net/http"
)

func GetImage(ct *gin.Context) {
	var request requests.GetImage
	err := ct.ShouldBind(&request)
	if err != nil {
		api.ReturnError(http.StatusBadRequest, err.Error(), ct)
		return
	}

	res, err := http.Get(request.Url)
	if err != nil {
		log.Printf("Failed to get image|%s|Error:%s", request.Url, err.Error())
		api.ReturnError(http.StatusBadRequest, "The image cannot found", ct)
		return
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)
	if err != nil {
		api.ReturnError(http.StatusBadRequest, "The image format not support yet", ct)
		return
	}
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Printf("Failed to encode image|%s|Error:%s", request.Url, err.Error())
		api.ReturnError(http.StatusInternalServerError, "Unable to encode image", ct)
		return
	}

	contextType := http.DetectContentType(buffer.Bytes())
	fileName := utils.GetOrCreateFileName(ct.Writer.Header().Get(consts.HeaderFileName), contextType)
	ct.Header(consts.HeaderFileName, fileName)

	api.ReturnFile(http.StatusOK, contextType, buffer.Bytes(), ct)
}
