package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"vitea/internal/database"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
)

func (s *Server) CreateResumeHandler(c echo.Context) error {
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

	resume := database.Resume{
		S3Key:  s3Key,
		UserID: 1,
	}

	repo := database.NewResumesRepository(s.db.DB())
	err = repo.Create(resume)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func (s *Server) GetResumesHandler(c echo.Context) error {
	repo := database.NewResumesRepository(s.db.DB())
	resumes, err := repo.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resumes)
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

func (s *Server) DeleteResumeHandler(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	repo := database.NewResumesRepository(s.db.DB())
	err = repo.Delete(idInt)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "OK")
}
