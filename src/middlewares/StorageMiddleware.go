package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"image-functions/src/consts"
	"image-functions/src/services/storage"
	"log"
	"net/http"
	"os"
)

type storageWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (sw storageWriter) Write(buf []byte) (int, error) {
	sw.body.Write(buf)
	return sw.ResponseWriter.Write(buf)
}

func StorageMiddleware() gin.HandlerFunc {
	return func(ct *gin.Context) {
		if os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
			return
		}
		var storageService storage.Storage = &storage.S3Service{
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			Region:          os.Getenv("AWS_DEFAULT_REGION"),
			Bucket:          os.Getenv("IMAGE_STORAGE_BUCKET"),
			Endpoint:        os.Getenv("AWS_ENDPOINT"),
			ForcePathStyle:  os.Getenv("S3_FORCE_PATH_STYLE") == "true",
		}
		err := storageService.Initial()
		if err != nil {
			log.Printf("StorageMiddleware|Failed to create storageService|Error%s", err)
			return
		}

		if !storageService.BucketExists(os.Getenv("IMAGE_STORAGE_BUCKET")) {
			log.Printf("StorageMiddleware|Failed to create storageService|No bucket: %s", os.Getenv("IMAGE_STORAGE_BUCKET"))
			return
		}

		// Note: Get the file from the storage and return the file directly if it is exists
		fileName, hasFileNameQuery := ct.GetQuery(consts.HeaderFileName)
		if hasFileNameQuery {
			ct.Header(consts.HeaderFileName, fileName)
			storageBuf, err := storageService.GetFile(fileName)
			if err == nil {
				ct.Data(http.StatusOK, http.DetectContentType(storageBuf), storageBuf)
				ct.Abort()
				return
			}
		}

		storageWriter := &storageWriter{body: bytes.NewBufferString(""), ResponseWriter: ct.Writer}
		ct.Writer = storageWriter
		ct.Next()

		// Note: Save the file into storage service
		fileName = ct.Writer.Header().Get(consts.HeaderFileName)
		if fileName != "" && ct.Writer.Status() < 300 && storageWriter.body != nil {
			defer storageService.UploadFile(storageWriter.body.Bytes(), fileName)
		}
	}
}
