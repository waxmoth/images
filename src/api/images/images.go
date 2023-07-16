package images

import (
	"github.com/gin-gonic/gin"
	"image-functions/src/api"
	"image-functions/src/consts"
	"image-functions/src/models/requests"
	"image-functions/src/utils"
	"log"
	"net/http"
)

// GetImage support fetch image from url and return to client
func GetImage(ct *gin.Context) {
	var request requests.GetImage
	err := ct.ShouldBind(&request)
	if err != nil {
		api.ReturnError(http.StatusBadRequest, err.Error(), ct)
		return
	}

	parsedURL, err := utils.ParseURL(request.URL, ct)
	if err != nil {
		log.Printf("Failed to parse url|%s|Error:%s", request.URL, err.Error())
		api.ReturnError(http.StatusBadRequest, "The url is invalid", ct)
		return
	}

	res, err := http.Get(parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.Path)
	if err != nil {
		log.Printf("Failed to get image|%s|Error:%s", request.URL, err.Error())
		api.ReturnError(http.StatusBadRequest, "The image cannot found", ct)
		return
	}
	defer res.Body.Close()

	img, format, err := utils.DecodeImage(res.Body)
	if err != nil {
		log.Printf("Failed decode image|%s|Error:%s", format, err.Error())
		api.ReturnError(http.StatusBadRequest, "The image format not support yet", ct)
		return
	}

	if request.Width > 0 && request.Height > 0 {
		switch request.Type {
		case consts.HandleResize:
			img = utils.ResizeImage(img, request.Width, request.Height)
		case consts.HandleCrop:
			img = utils.CropImage(img, request.PositionX, request.PositionY, request.Width, request.Height)
		default:
			img = utils.CropImage(img, request.PositionX, request.PositionY, request.Width, request.Height)
		}
	}

	buffer, err := utils.ImageToBuffer(img, format)
	if err != nil {
		log.Printf("Failed to encode image|%s|Error:%s", request.URL, err.Error())
		api.ReturnError(http.StatusInternalServerError, "Unable to encode image", ct)
		return
	}

	contextType := http.DetectContentType(buffer.Bytes())
	fileName := utils.GetOrCreateFileName(ct.Writer.Header().Get(consts.HeaderFileName), contextType)
	ct.Header(consts.HeaderFileName, fileName)

	api.ReturnFile(http.StatusOK, contextType, buffer.Bytes(), ct)
}
