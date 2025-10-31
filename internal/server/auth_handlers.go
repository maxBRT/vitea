package server

import (
	"time"
	"vitea/internal/auth"
	"vitea/internal/database/sqlc"

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

	user, err := s.db.Queries().FindUserByEmail(c.Request().Context(), sqlc.FindUserByEmailParams{
		Email: req.Email,
	})
	if err != nil {
		return err
	}

	if err := auth.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return err
	}

	token, err := auth.GenerateToken(user.ID, s.jwtSecret, time.Minute*5)
	if err != nil {
		return err
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		return err
	}

	_, err = s.db.Queries().CreateRefreshToken(c.Request().Context(), sqlc.CreateRefreshTokenParams{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	})
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{
		"access_token":  token,
		"refresh_token": refreshToken,
	})

}
