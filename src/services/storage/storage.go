package storage

// Storage interface
type Storage interface {
	Initial() error
	UploadFile(buf []byte, fileName string) bool
	GetFile(fileName string) ([]byte, error)
	BucketExists(bucketName string) bool
	FileExists(fileName string) bool
}
