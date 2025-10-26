package server

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"vitea/internal/database"
)

type Server struct {
	port     int
	s3Bucket string
	s3Region string
	s3Client *s3.Client
	db       database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	region := os.Getenv("S3_REGION")
	if region == "" {
		log.Fatal("S3_REGION is not set")
	}
	bucket := os.Getenv("S3_BUCKET")
	if bucket == "" {
		log.Fatal("S3_BUCKET is not set")
	}
	config, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(os.Getenv("S3_REGION")))
	if err != nil {
		log.Fatal(err)
	}

	NewServer := &Server{
		port:     port,
		s3Bucket: bucket,
		s3Region: region,
		s3Client: s3.NewFromConfig(config),
		db:       database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:        fmt.Sprintf(":%d", NewServer.port),
		Handler:     NewServer.RegisterRoutes(),
		IdleTimeout: time.Minute,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
