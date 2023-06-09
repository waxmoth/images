package storage

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log"
)

type S3Service struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	Bucket          string
	Endpoint        string
	ForcePathStyle  bool
	s3Client        *s3.S3
}

func (storageServ *S3Service) Initial() error {
	credentials := credentials.NewStaticCredentials(storageServ.AccessKeyID, storageServ.SecretAccessKey, "")
	_, err := credentials.Get()
	if err != nil {
		log.Printf("S3Service|Initial|Bad credentials|Error: %s", err.Error())
		return err
	}
	cfg := aws.Config{Endpoint: &storageServ.Endpoint, Region: &storageServ.Region, Credentials: credentials}
	cfg.WithS3ForcePathStyle(storageServ.ForcePathStyle)
	sess, err := session.NewSession(&cfg)
	if err != nil {
		log.Printf("S3Service|Initial|Cannot create session|Error: %s", err.Error())
		return err
	}
	storageServ.s3Client = s3.New(sess)
	return nil
}

func (storageServ *S3Service) GetFile(fileName string) ([]byte, error) {
	res, err := storageServ.s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(storageServ.Bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	return buf, err
}

func (storageServ *S3Service) UploadFile(buf []byte, fileName string) bool {
	_, err := storageServ.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(storageServ.Bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buf),
	})
	if err != nil {
		log.Printf("S3Service|UploadFile|Cannot upload the file %s|Error: %s", fileName, err)
	}
	return err == nil
}

func (storageServ *S3Service) BucketExists(bucketName string) bool {
	_, err := storageServ.s3Client.HeadBucket(&s3.HeadBucketInput{Bucket: aws.String(bucketName)})
	if err != nil {
		log.Printf("S3Service|BucketExists|Cannot get bucket %s|Error:%s", bucketName, err.Error())
		return false
	}

	return true
}

func (storageServ *S3Service) FileExists(fileName string) bool {
	_, err := storageServ.s3Client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(storageServ.Bucket),
		Key:    aws.String(fileName),
	})
	return err == nil
}
