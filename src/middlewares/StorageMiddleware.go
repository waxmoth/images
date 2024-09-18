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

// StorageMiddleware handle return or upload file to the storage service
func StorageMiddleware() gin.HandlerFunc {
	return func(ct *gin.Context) {
		var storageService = getStorageService(os.Getenv("STORAGE_TYPE"))
		if storageService == nil {
			return
		}
		err := storageService.Initial()
		if err != nil {
			log.Printf("StorageMiddleware|Failed to create storageService|Error: %s", err)
			return
		}

		if !storageService.BucketExists(os.Getenv("IMAGE_STORAGE_BUCKET")) {
			log.Printf("StorageMiddleware|Failed to create storageService|No bucket: %s", os.Getenv("IMAGE_STORAGE_BUCKET"))
			return
		}

		// Note: Get the file from the storage and return the file directly if it is existing
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
		if fileName != "" && ct.Writer.Status() < 300 {
			if fileData, exists := ct.Get(consts.FileData); exists && fileData != nil {
				defer storageService.UploadFile(fileData.([]byte), fileName)
				return
			}
			if storageWriter.body.Len() > 10 {
				defer storageService.UploadFile(storageWriter.body.Bytes(), fileName)
				return
			}
		}
	}
}

func getStorageService(storageType string) storage.Storage {
	if os.Getenv("IMAGE_STORAGE_BUCKET") == "" {
		return nil
	}
	switch storageType {
	case "local":
		return &storage.LocalService{
			LocalFolder: os.Getenv("IMAGE_STORAGE_BUCKET"),
		}
	case "s3":
		if os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
			return nil
		}
		return &storage.S3Service{
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			Region:          os.Getenv("AWS_DEFAULT_REGION"),
			Bucket:          os.Getenv("IMAGE_STORAGE_BUCKET"),
			Endpoint:        os.Getenv("AWS_ENDPOINT"),
			ForcePathStyle:  os.Getenv("S3_FORCE_PATH_STYLE") == "true",
		}
	case "gcs":
		if os.Getenv("GOOGLE_CLOUD_CREDENTIALS") == "" {
			return nil
		}
		return &storage.GCSService{
			Credits:  os.Getenv("GOOGLE_CLOUD_CREDENTIALS"),
			Bucket:   os.Getenv("IMAGE_STORAGE_BUCKET"),
			Endpoint: os.Getenv("GOOGLE_STORAGE_EMULATOR_HOST"),
		}
	default:
		return nil
	}
}
