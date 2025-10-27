package database

import "database/sql"

type User struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
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
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (r *UserRepository) Get(id int) (User, error) {
	var user User
	err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err
}

func (r *UserRepository) Create(user User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (first_name, last_name, email, hashed_password) VALUES ($1, $2, $3, $4)",
		user.FirstName,
		user.LastName,
		user.Email,
		user.HashedPassword,
	)
	return err
}

func (r *UserRepository) Update(user User) error {
	return nil
}

func (r *UserRepository) Delete(id int) error {
	err := r.db.QueryRow("DELETE FROM users WHERE id = $1", id).Scan()
	return err
}
