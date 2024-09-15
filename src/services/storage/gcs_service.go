package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"io"
	"log"
	"os"
)

// GCSService handle files in Google Cloud Storage
type GCSService struct {
	Credits   string
	Bucket    string
	Endpoint  string
	ctx       context.Context
	gcsClient *storage.Client
}

// Initial initialize GCS service
func (storageServ *GCSService) Initial() error {
	opts := []option.ClientOption{
		option.WithCredentialsJSON([]byte(storageServ.Credits)),
	}
	if storageServ.Endpoint != "" {
		err := os.Setenv("STORAGE_EMULATOR_HOST", storageServ.Endpoint)
		if err != nil {
			return err
		}
	}

	storageServ.ctx = context.Background()
	gcsClient, err := storage.NewClient(storageServ.ctx, opts...)
	if err != nil {
		return err
	}
	storageServ.gcsClient = gcsClient
	return nil
}

// GetFile get file buffer from Google Cloud Storage, return error if cannot get it
func (storageServ *GCSService) GetFile(fileName string) ([]byte, error) {
	res, err := storageServ.gcsClient.Bucket(storageServ.Bucket).Object(fileName).NewReader(storageServ.ctx)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return io.ReadAll(res)
}

// UploadFile upload the file into Google Cloud Storage service
func (storageServ *GCSService) UploadFile(buf []byte, fileName string) bool {
	writer := storageServ.gcsClient.Bucket(storageServ.Bucket).Object(fileName).NewWriter(storageServ.ctx)
	defer writer.Close()
	_, err := writer.Write(buf)
	if err != nil {
		log.Printf("GCSService|UploadFile|Cannot upload the file %s|Error: %s", fileName, err)
	}
	return err == nil
}

// BucketExists check if the bucket exists
func (storageServ *GCSService) BucketExists(bucketName string) bool {
	_, err := storageServ.gcsClient.Bucket(bucketName).Attrs(storageServ.ctx)
	if err != nil {
		log.Printf("GCSService|BucketExists|Cannot get bucket %s|Error: %s", bucketName, err.Error())
		return false
	}

	return true
}

// FileExists check if the file exists
func (storageServ *GCSService) FileExists(fileName string) bool {
	_, err := storageServ.gcsClient.Bucket(storageServ.Bucket).Object(fileName).Attrs(storageServ.ctx)
	return err == nil
}
