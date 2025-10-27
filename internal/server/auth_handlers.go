package server

import (
	"time"
	"vitea/internal/auth"
	"vitea/internal/database"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) LoginHandler(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	repo := database.NewUserRepository(s.db.DB())
	user, err := repo.FindByEmail(req.Email)
	if err != nil {
		return err
	}
	if err := auth.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return err
	}

	token, err := auth.GenerateToken(user.ID, s.jwtSecret, time.Minute*15)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{
		"access_token": token,
	})

}
