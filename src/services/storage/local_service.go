package storage

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const safeDir = "/tmp/"

// LocalService handle files in local storage
type LocalService struct {
	LocalFolder string
}

// Initial create local storage service from the configures
func (storageServ *LocalService) Initial() error {
	if !storageServ.BucketExists(storageServ.LocalFolder) {
		log.Printf("LocalService|Initial|No bucket: %s", storageServ.LocalFolder)
		resolvedFolder, err := storageServ.resolveFileName(storageServ.LocalFolder)
		if err != nil {
			return err
		}
		err = os.MkdirAll(resolvedFolder, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetFile get file buffer from local storage, return error if you cannot get it
func (storageServ *LocalService) GetFile(fileName string) ([]byte, error) {
	resolvedFileName, err := storageServ.resolveFileName(storageServ.LocalFolder + "/" + fileName)
	if err != nil {
		return nil, err
	}
	content, err := os.ReadFile(resolvedFileName)
	if err != nil {
		return nil, err
	}
	return content, err
}

// UploadFile upload the file into local storage service
func (storageServ *LocalService) UploadFile(buf []byte, fileName string) bool {
	resolvedFileName, err := storageServ.resolveFileName(storageServ.LocalFolder + "/" + fileName)
	if resolvedFileName == "" || err != nil {
		return false
	}
	err = os.WriteFile(resolvedFileName, buf, 0644)
	if err != nil {
		log.Printf("LocalService|UploadFile|Cannot upload the file %s|Error: %s", fileName, err)
	}
	return err == nil
}

// BucketExists check if the bucket exists
func (storageServ *LocalService) BucketExists(folder string) bool {
	resolvedFolder, err := storageServ.resolveFileName(folder)
	if err != nil {
		return false
	}
	if info, err := os.Stat(resolvedFolder); os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}

// FileExists check if the file exists
func (storageServ *LocalService) FileExists(fileName string) bool {
	resolvedFileName, err := storageServ.resolveFileName(storageServ.LocalFolder + "/" + fileName)
	if err != nil {
		return false
	}

	_, err = os.Stat(resolvedFileName)
	return err == nil
}

// resolveFileName resolve the file name, in case the path injection issue
func (storageServ *LocalService) resolveFileName(fileName string) (string, error) {
	resolvedFileName, err := filepath.Abs(filepath.Join(safeDir, fileName))
	if err != nil || !strings.HasPrefix(resolvedFileName, safeDir) {
		return "", errors.New("invalid file name")
	}

	return resolvedFileName, nil
}
