package server

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"vitea/internal/auth"
	"vitea/internal/database"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/resend/resend-go/v2"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

func (s *Server) CreateUserHandler(c echo.Context) error {
	user := User{}
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
	usr := database.User{
		ID:             uuid.New(),
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		HashedPassword: user.Password,
	}

	repo := database.NewUserRepository(s.db.DB())
	if err := repo.Create(&usr); err != nil {
		return err
	}

	fmt.Println(usr.ID)
	token, err := auth.GenerateToken(usr.ID, s.jwtSecret, time.Hour*24)
	if err != nil {
		return err
	}

	link := fmt.Sprintf("http://%s/api/verify/%s", s.baseURL, token)

	if err := sendVerificationEmail(user.Email, link); err != nil {
		return err
	}

	repo.VerifyLink(user.ID, token)

	return c.NoContent(http.StatusCreated)
}

func (s *Server) GetUserHandler(c echo.Context) error {
	id := c.Get("user_id")

	repo := database.NewUserRepository(s.db.DB())
	user, err := repo.Get(id.(uuid.UUID))
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

	repo := database.NewUserRepository(s.db.DB())
	if err := repo.Delete(id); err != nil {
		return err
	}
	return c.String(http.StatusOK, "OK")
}

func sendVerificationEmail(to, link string) error {
	client := resend.NewClient(os.Getenv("RESEND_API_KEY"))
	fmt.Println(link)
	html := fmt.Sprintf(`
        <h2>Welcome to Vitae!</h2>
        <p>To complete your registration, click below:</p>
        <a href="%s" style="background:#2563eb;color:white;padding:12px 24px;border-radius:8px;text-decoration:none;">Verify My Email</a>
    `, link)
	_, err := client.Emails.Send(&resend.SendEmailRequest{
		From:    "Acme <onboarding@resend.dev>",
		To:      []string{to},
		Subject: "Welcome to Vitae!",
		Html:    html,
	})
	return err
}
