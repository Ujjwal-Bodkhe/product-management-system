package service

import (
	"log"
	//"github.com/Ujjwal-Bodkhe/product-management-system/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ImageProcessor struct {
	S3Client *s3.Client
}

func NewImageProcessor() *ImageProcessor {
	creds := credentials.NewStaticCredentialsProvider("your-access-key", "your-secret-key", "")
	cfg := aws.Config{
		Region:      "us-west-2",
		Credentials: creds,
	}

	s3Client := s3.NewFromConfig(cfg)
	return &ImageProcessor{S3Client: s3Client}
}

func (p *ImageProcessor) ProcessImage(imageURL string) (string, error) {
	// Here you would download the image, compress it and upload to S3
	// For simplicity, we'll assume the image is processed and uploaded
	// Replace with actual S3 interaction code as needed

	log.Printf("Processing image: %s", imageURL)
	return "https://your-s3-bucket-url/image.jpg", nil
}
