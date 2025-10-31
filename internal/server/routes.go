package server

import (
	"net/http"
	"vitea/internal/auth"
	"vitea/internal/database/sqlc"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	api := e.Group("/api")

	api.GET("/verify/:token", s.VerifyHandler)
	api.POST("/refresh", s.RefreshHandler)
	api.POST("/login", s.LoginHandler)
	api.POST("/users", s.CreateUserHandler)
	api.GET("/users", s.ListUsersHandler)
	api.GET("/resumes", s.GetResumesHandler)
	api.GET("/health", s.healthHandler)
	api.DELETE("/users/:id", s.DeleteUserHandler)

	secure := api.Group("")
	secure.Use(auth.JWTMiddleware(s.jwtSecret))

	secure.GET("/me", s.GetUserHandler)
	secure.POST("/resumes", s.CreateResumeHandler)
	secure.DELETE("/resumes/:id", s.DeleteResumeHandler)

	return e
}

func (s *Server) VerifyHandler(c echo.Context) error {
	tokenString := c.Param("token")
	userID, err := auth.ValidateJWT(tokenString, s.jwtSecret)
	if err != nil {
		return err
	}
	if _, err := s.db.Queries().VerifyUser(c.Request().Context(), sqlc.VerifyUserParams{
		ID: userID,
	}); err != nil {
		return err
	}
	if err := s.db.Queries().DeleteVerifyLink(c.Request().Context(), sqlc.DeleteVerifyLinkParams{
		UserID: userID,
	}); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, s.frontendURL)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
