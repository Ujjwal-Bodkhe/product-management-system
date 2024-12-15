package service

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nfnt/resize"
)

// compressAndUploadImages compresses the product images and uploads them to S3.
func compressAndUploadImages(imageURLs []string, s3Bucket string) ([]string, error) {
	var compressedURLs []string

	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	s3Client := s3.New(sess)

	for _, imageURL := range imageURLs {
		// Download the image
		resp, err := http.Get(imageURL)
		if err != nil {
			log.Printf("Failed to download image: %v", err)
			continue
		}
		defer resp.Body.Close()

		// Decode the image
		img, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Printf("Failed to decode image: %v", err)
			continue
		}

		// Resize the image
		resizedImg := resize.Resize(800, 0, img, resize.Lanczos3)

		// Compress the image to JPEG
		var buf bytes.Buffer
		err = jpeg.Encode(&buf, resizedImg, nil)
		if err != nil {
			log.Printf("Failed to encode compressed image: %v", err)
			continue
		}

		// Upload the compressed image to S3
		fileName := fmt.Sprintf("compressed/%d.jpg", time.Now().UnixNano())
		_, err = s3Client.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(s3Bucket),
			Key:    aws.String(fileName),
			Body:   bytes.NewReader(buf.Bytes()),
			ACL:    aws.String("public-read"),
		})
		if err != nil {
			log.Printf("Failed to upload image to S3: %v", err)
			continue
		}

		// Add the S3 URL of the uploaded image to the list of compressed URLs
		compressedURLs = append(compressedURLs, fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s3Bucket, fileName))
	}

	return compressedURLs, nil
}
