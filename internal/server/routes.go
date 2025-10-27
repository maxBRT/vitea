package server

import (
	"net/http"
	"vitea/internal/auth"

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
	authRequired := e.Group("")
	authRequired.Use(auth.JWTMiddleware(s.jwtSecret))

	authRequired.GET("/users/:id", s.GetUserHandler)
	authRequired.DELETE("/users/:id", s.DeleteUserHandler)
	authRequired.POST("/resumes", s.CreateResumeHandler)
	authRequired.DELETE("/resumes/:id", s.DeleteResumeHandler)

	e.POST("/auth/login", s.LoginHandler)
	e.POST("/users", s.CreateUserHandler)
	e.GET("/users", s.ListUsersHandler)
	e.GET("/resumes", s.GetResumesHandler)

	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
