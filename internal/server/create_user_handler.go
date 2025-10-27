package server

import (
	"net/http"
	"vitea/internal/database"

	"github.com/labstack/echo/v4"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (s *Server) CreateUserHandler(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	repo := database.NewUserRepository(s.db.DB())
	err := repo.Create(database.User{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		HashedPassword: "password",
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
