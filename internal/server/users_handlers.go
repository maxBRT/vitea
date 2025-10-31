package server

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"vitea/internal/auth"
	"vitea/internal/database/sqlc"
)

type UserRequest struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func (s *Server) CreateUserHandler(c echo.Context) error {
	user := UserRequest{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if user.FirstName == "" || user.LastName == "" || user.Email == "" || user.Password == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	hash, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash

	usr, err := s.db.Queries().CreateUser(c.Request().Context(), sqlc.CreateUserParams{
		ID:             uuid.New(),
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		HashedPassword: user.Password,
	})
	if err != nil {
		return err
	}

	token, err := auth.GenerateToken(usr.ID, s.jwtSecret, time.Hour*24)
	if err != nil {
		return err
	}

	link := fmt.Sprintf("http://%s/api/verify/%s", s.baseURL, token)

	if err := auth.SendVerificationEmail(user.Email, link); err != nil {
		return err
	}

	if err := s.db.Queries().CreateVerifyLink(c.Request().Context(), sqlc.CreateVerifyLinkParams{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}); err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func (s *Server) GetUserHandler(c echo.Context) error {
	id := c.Get("user_id")

	user, err := s.db.Queries().GetUser(c.Request().Context(), sqlc.GetUserParams{
		ID: id.(uuid.UUID),
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (s *Server) ListUsersHandler(c echo.Context) error {
	users, err := s.db.Queries().GetAllUsers(c.Request().Context())
	if err != nil {
		return err
	}
	res := []UserResponse{}
	for _, u := range users {
		res = append(res, UserResponse{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			CreatedAt: u.CreatedAt.Format(time.RFC3339),
			UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
		})
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Server) DeleteUserHandler(c echo.Context) error {
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err := s.db.Queries().DeleteUser(c.Request().Context(), sqlc.DeleteUserParams{
		ID: userID,
	}); err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}
