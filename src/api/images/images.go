package images

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"image"
	"image-functions/src/api"
	"image-functions/src/models/requests"
	"image/jpeg"
	"log"
	"net/http"
)

func GetImage(ct *gin.Context) {
	var request requests.GetImage
	err := ct.Bind(&request)
	if err != nil {
		api.ReturnError(http.StatusBadRequest, err.Error(), ct)
		return
	}
	res, err := http.Get(request.Url)
	defer res.Body.Close()
	if err != nil {
		log.Printf("Failed to get image|%s|Error:%s", request.Url, err.Error())
		api.ReturnError(http.StatusBadRequest, "The image cannot found", ct)
		return
	}

	img, _, err := image.Decode(res.Body)
	if err != nil {
		api.ReturnError(http.StatusBadRequest, "The image format not support yet", ct)
		return
	}
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		api.ReturnError(http.StatusInternalServerError, "Unable to encode image", ct)
	}

	api.ReturnFile(http.StatusOK, "image/jpeg", buffer.Bytes(), ct)
}
