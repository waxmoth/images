package images

import (
	"github.com/gin-gonic/gin"
	"image-functions/src/api"
	"image-functions/src/consts"
	"image-functions/src/models/requests"
	"image-functions/src/models/responses"
	"image-functions/src/utils"
	"log"
	"net/http"
)

// GetImage support fetch image from url and return to client
//
//	@Router			/api/image [get]
//	@Summary		Get image
//	@Schemes		http https
//	@Description	Fetch the image from url and return to client
//	@Description	You can resize the image by query `width` and `height`
//	@Tags			image
//	@Accept			application/json
//	@Produce		jpeg
//	@Param			Authorization	header		string				true	"Bearer token"
//	@Param			object			query		requests.GetImage	false	"Get image request payload"
//	@Success		200				{string}	image				"The image file"
//	@Header			200				{string}	File-Name			"The cached image file name"
//	@Failure		400				object		api.Error
//	@Failure		401				{string}	string	"Unauthorized"
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

// UploadImage support upload image to s3
//
//	@Router			/api/image [post]
//	@Summary		Upload image
//	@Schemes		http https
//	@Description	Upload the image to service
//	@Tags			image
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			Authorization	header		string					true	"Bearer token"
//	@Param			object			body		requests.UploadImage	true	"Upload image request payload"
//	@Param			file			formData	file					true	"The image file"
//	@Success		200				object		api.SuccessResponse{data=responses.UploadImage}
//	@Failure		400				object		api.Error
//	@Failure		401				{string}	string	"Unauthorized"
func UploadImage(ct *gin.Context) {
	var request requests.UploadImage
	err := ct.ShouldBind(&request)
	if err != nil {
		api.ReturnError(http.StatusBadRequest, err.Error(), ct)
		return
	}
	file, err := request.File.Open()
	if err != nil {
		api.ReturnError(http.StatusBadRequest, "Cannot open the file", ct)
		return
	}
	defer file.Close()
	img, format, err := utils.DecodeImage(file)
	if err != nil {
		log.Printf("Failed decode image|%s|Error:%s", format, err.Error())
		api.ReturnError(http.StatusBadRequest, "The image format not support yet", ct)
		return
	}
	buffer, err := utils.ImageToBuffer(img, format)
	if err != nil {
		log.Printf("Failed to encode image|Error:%s", err.Error())
		api.ReturnError(http.StatusInternalServerError, "Unable to encode image", ct)
		return
	}

	fileName := request.Name
	if fileName == "" {
		contextType := http.DetectContentType(buffer.Bytes())
		fileName = utils.GetOrCreateFileName(request.Name, contextType)
	}

	ct.Header(consts.HeaderFileName, fileName)
	ct.Set(consts.FileData, buffer.Bytes())

	api.ReturnSuccess(http.StatusOK, responses.UploadImage{Name: fileName}, "Success", ct)
}
