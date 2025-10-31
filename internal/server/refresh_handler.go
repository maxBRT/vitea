package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"vitea/internal/auth"
	"vitea/internal/database/sqlc"
)

func (s *Server) RefreshHandler(c echo.Context) error {
	refreshToken, err := auth.GetBearerToken(c.Request().Header)
	if err != nil {
		return err
	}
	token, err := s.db.Queries().GetRefreshTokenByToken(c.Request().Context(), sqlc.GetRefreshTokenByTokenParams{
		Token: refreshToken,
	})
	if err != nil {
		return err
	}

	if token.ExpiresAt.Before(time.Now()) {
		return c.NoContent(http.StatusUnauthorized)
	}

	jwtToken, err := auth.GenerateToken(token.UserID, s.jwtSecret, time.Hour*24)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token": jwtToken,
	})
}
