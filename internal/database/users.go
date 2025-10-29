package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
	Verified       bool      `json:"verified"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Verified,
	)
	return user, err
}

func (r *UserRepository) GetAll() ([]User, error) {
	var users []User
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.HashedPassword,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Verified,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (r *UserRepository) Get(id uuid.UUID) (User, error) {
	var user User
	err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Verified,
	)
	return user, err
}

func (r *UserRepository) Create(user *User) error {
	query := `
        INSERT INTO users (id, first_name, last_name, email, hashed_password)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, first_name, last_name, email, hashed_password, verified, created_at;
    `

	return r.db.QueryRow(
		query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.HashedPassword,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.HashedPassword,
		&user.Verified,
		&user.CreatedAt,
	)
}

func (r *UserRepository) Update(user User) error {
	return nil
}

func (r *UserRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (r *UserRepository) VerifyLink(user_id uuid.UUID, token string) error {
	_, err := r.db.Exec(
		"INSERT INTO user_verify_link (user_id, token, expires_at) VALUES ($1, $2, $3)",
		user_id,
		token,
		time.Now().Add(time.Hour*24),
	)
	return err
}

func (r *UserRepository) DeleteVerifyLink(user_id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM user_verify_link WHERE user_id = $1", user_id)
	return err
}

func (r *UserRepository) Verify(id uuid.UUID) error {
	fmt.Println("verifying user")

	result, err := r.db.Exec(`UPDATE users SET verified = TRUE WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get affected rows: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("no user found with id %s", id)
	}

	fmt.Printf("✅ User %s verified (%d row updated)\n", id, rows)
	return nil
}
