package storage

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client *s3.Client
	bucket string
}

// NewS3Client initializes a new S3 client
func NewS3Client() (*S3Client, error) {
	// Load AWS configuration from environment variables or default config
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("S3_REGION")),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("S3_ACCESS_KEY_ID"),
			os.Getenv("S3_SECRET_ACCESS_KEY"),
			"",
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	// Initialize S3 client
	s3Client := s3.NewFromConfig(cfg)

	return &S3Client{
		client: s3Client,
		bucket: os.Getenv("S3_BUCKET_NAME"),
	}, nil
}

// UploadImage uploads an image to S3 and returns the S3 object URL
func (s *S3Client) UploadImage(fileName string, fileContent []byte) (string, error) {
	// Create an Upload input
	input := &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &fileName,
		Body:   bytes.NewReader(fileContent),
		ContentType: aws.String("image/jpeg"), // or determine MIME type based on file
	}

	// Upload the file to S3
	_, err := s.client.PutObject(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("unable to upload image %v", err)
	}

	// Construct and return the file URL
	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, os.Getenv("S3_REGION"), fileName)
	return fileURL, nil
}

// DownloadImage retrieves an image from S3
func (s *S3Client) DownloadImage(fileName string) ([]byte, error) {
	// Create a GetObject input
	input := &s3.GetObjectInput{
		Bucket: &s.bucket,
		Key:    &fileName,
	}

	// Retrieve the image from S3
	resp, err := s.client.GetObject(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("unable to download image %v", err)
	}
	defer resp.Body.Close()

	// Read the body of the response
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read image data %v", err)
	}

	return buf.Bytes(), nil
}

//CompressAndUploadImage compresses the image (could be a placeholder for actual compression) and uploads to S3
func (s *S3Client) CompressAndUploadImage(fileName string, imageData []byte) (string, error) {
	// Placeholder for image compression logic
	// In real implementation, you could use a package like `image/jpeg` or `image/png` to compress the image
	compressedImage := imageData // This should be replaced with actual image compression logic

	// Upload the compressed image to S3
	return s.UploadImage(fileName, compressedImage)
}
