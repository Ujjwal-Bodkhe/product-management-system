package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type S3Client struct {
	s3Client *s3.S3
	bucket   string
}

func NewS3Client(region, bucket string) *S3Client {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Fatal("Unable to initialize AWS session: ", err)
	}
	return &S3Client{
		s3Client: s3.New(sess),
		bucket:   bucket,
	}
}

func (s *S3Client) UploadImage(key string, filePath string) error {
	// Upload the image to S3
	_, err := s.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   aws.String(filePath),
	})
	return err
}
