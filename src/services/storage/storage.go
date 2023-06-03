package storage

type Storage interface {
	Initial() error
	UploadFile(buf []byte, fileName string) (string, error)
	GetFile(fileName string) ([]byte, error)
	BucketExists(bucketName string) bool
	FileExists(fileName string) bool
}
