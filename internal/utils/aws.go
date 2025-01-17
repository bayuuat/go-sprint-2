package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// CreateSession creates a new AWS session
func CreateSession() (*session.Session, error) {
	region := os.Getenv("AWS_REGION")

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	return session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKey,
			secretKey,
			"",
		),
	})
}

// UploadFileToS3 uploads a file to S3 with a given prefix
func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader, bucket, prefix string) (string, error) {
	defer file.Close()

	sess, err := CreateSession()
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	timestamp := time.Now().Format("20060102150405")
	objectKey := fmt.Sprintf("%s/%s_%s", prefix, timestamp, fileHeader.Filename)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(objectKey),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
	})

	region := os.Getenv("AWS_REGION")
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, objectKey)

	return url, err
}
