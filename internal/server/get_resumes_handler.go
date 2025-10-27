package server

import (
	"net/http"
	"vitea/internal/database"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetResumesHandler(c echo.Context) error {
	repo := database.NewResumesRepository(s.db.DB())
	resumes, err := repo.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resumes)
}
