package storage

import (
	"log"
	"os"
)

// LocalService handle files in local
type LocalService struct {
	LocalFolder string
}

// Initial create s3 service from the configures
func (storageServ *LocalService) Initial() error {
	if !storageServ.BucketExists(storageServ.LocalFolder) {
		log.Printf("LocalService|Initial|No bucket: %s", storageServ.LocalFolder)
		err := os.MkdirAll(storageServ.LocalFolder, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetFile get file buffer from s3, return error if you cannot get it
func (storageServ *LocalService) GetFile(fileName string) ([]byte, error) {
	content, err := os.ReadFile(storageServ.LocalFolder + "/" + fileName)
	if err != nil {
		return nil, err
	}
	return content, err
}

// UploadFile upload the file into s3 service
func (storageServ *LocalService) UploadFile(buf []byte, fileName string) bool {
	err := os.WriteFile(storageServ.LocalFolder+"/"+fileName, buf, 0644)
	if err != nil {
		log.Printf("LocalService|UploadFile|Cannot upload the file %s|Error: %s", fileName, err)
	}
	return err == nil
}

// BucketExists check if the bucket exists
func (storageServ *LocalService) BucketExists(folder string) bool {
	if info, err := os.Stat(folder); os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}

// FileExists check if the file exists
func (storageServ *LocalService) FileExists(fileName string) bool {
	_, err := os.Stat(storageServ.LocalFolder + "/" + fileName)
	return err == nil
}
