package server

import (
	"net/http"
	"strconv"
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

	return c.NoContent(http.StatusCreated)
}

func (s *Server) GetUserHandler(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	repo := database.NewUserRepository(s.db.DB())
	user, err := repo.Get(idInt)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (s *Server) ListUsersHandler(c echo.Context) error {
	repo := database.NewUserRepository(s.db.DB())
	users, err := repo.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (s *Server) DeleteUserHandler(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	repo := database.NewUserRepository(s.db.DB())
	err = repo.Delete(idInt)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "OK")
}
