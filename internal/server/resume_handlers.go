package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"vitea/internal/database/sqlc"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
)

func (s *Server) CreateResumeHandler(c echo.Context) error {
	idVal := c.Get("user_id")
	if idVal == nil {
		return c.NoContent(http.StatusBadRequest)
	}
	idStr, ok := idVal.(string)
	if !ok {
		return c.NoContent(http.StatusBadRequest)
	}

	userID, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}

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

	_, err = s.db.Queries().CreateResume(c.Request().Context(), sqlc.CreateResumeParams{
		S3Key:  s3Key,
		UserID: userID,
	})
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}

func (s *Server) GetResumesHandler(c echo.Context) error {
	resumes, err := s.db.Queries().GetAllResumes(c.Request().Context())
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
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	s.db.Queries().DeleteResume(c.Request().Context(), sqlc.DeleteResumeParams{
		ID: int32(intID),
	})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "OK")
}
