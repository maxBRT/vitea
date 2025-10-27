package server

import (
	"net/http"
	"vitea/internal/database"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetUsersHandler(c echo.Context) error {
	repo := database.NewUserRepository(s.db.DB())
	users, err := repo.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
