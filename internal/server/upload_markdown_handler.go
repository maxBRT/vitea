package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
)

func (s *Server) UploadMarkdownHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	content_type := file.Header.Get("Content-Type")

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	buffer, err := extractHTML(src)
	if err != nil {
		return err
	}

	// Fucntion that writes the css and html to a file and returns the path

	// Upload the file to S3
	s3Key := fmt.Sprintf("resume/%s.html", uuid.NewString())
	obj := s3.PutObjectInput{
		Bucket:      &s.s3Bucket,
		Key:         &s3Key,
		Body:        bytes.NewReader(buffer),
		ContentType: &content_type,
	}
	s.s3Client.PutObject(context.Background(), &obj)

	return c.String(http.StatusOK, string(buffer))
}

func extractHTML(file io.Reader) ([]byte, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	md := goldmark.New()

	if err := md.Convert(content, buffer); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
