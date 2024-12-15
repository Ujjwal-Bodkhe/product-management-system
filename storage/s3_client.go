package storage

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/yourusername/product-management-system/config"
	"log"
)

func UploadToS3(fileName string, fileData []byte) (string, error) {
	// Load config
	cfg := config.LoadConfig()

	// Initialize an S3 session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: aws.NewStaticCredentials(cfg.AWSAccessKey, cfg.AWSSecretKey, ""),
	})
	if err != nil {
		log.Fatal("Failed to create session:", err)
		return "", err
	}

	s3Client := s3.New(sess)

	// Upload the file
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(cfg.AWSBucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(fileData),
	})
	if err != nil {
		log.Fatal("Failed to upload to S3:", err)
		return "", err
	}

	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", cfg.AWSBucketName, fileName)
	return fileURL, nil
}
