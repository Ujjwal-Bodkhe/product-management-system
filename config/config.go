package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	RedisURL      string
	RabbitMQURL   string
	AWSAccessKey  string
	AWSSecretKey  string
	AWSBucketName string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		RedisURL:      os.Getenv("REDIS_URL"),
		RabbitMQURL:   os.Getenv("RABBITMQ_URL"),
		AWSAccessKey:  os.Getenv("AWS_ACCESS_KEY"),
		AWSSecretKey:  os.Getenv("AWS_SECRET_KEY"),
		AWSBucketName: os.Getenv("AWS_S3_BUCKET"),
	}
}
