package database

import (
	"database/sql"
	"time"
)

type Resume struct {
	ID        int       `json:"id" db:"id"`
	S3Key     string    `json:"s3_key" db:"s3_key"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UserID    int       `json:"user_id" db:"user_id"`
}

type ResumeRepository struct {
	db *sql.DB
}

func NewResumesRepository(db *sql.DB) *ResumeRepository {
	return &ResumeRepository{db: db}
}

func (r *ResumeRepository) GetAll() ([]Resume, error) {
	rows, err := r.db.Query("SELECT id, s3_key, created_at, updated_at, user_id FROM resumes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resumes []Resume
	for rows.Next() {
		var resume Resume
		if err := rows.Scan(
			&resume.ID,
			&resume.S3Key,
			&resume.CreatedAt,
			&resume.UpdatedAt,
			&resume.UserID,
		); err != nil {
			return nil, err
		}
		resumes = append(resumes, resume)
	}
	return resumes, nil
}

func (r *ResumeRepository) Get(id int) (Resume, error) {
	var resume Resume
	err := r.db.QueryRow("SELECT * FROM resumes WHERE id = $1", id).Scan(
		&resume.ID,
		&resume.S3Key,
		&resume.CreatedAt,
		&resume.UpdatedAt,
		&resume.UserID,
	)
	return resume, err
}

func (r *ResumeRepository) Create(resume Resume) error {
	_, err := r.db.Exec("INSERT INTO resumes (s3_key, user_id) VALUES ($1, $2)", resume.S3Key, resume.UserID)
	return err
}

func (r *ResumeRepository) Update(resume Resume) error {
	_, err := r.db.Exec("UPDATE resumes SET s3_key = $1, created_at = $2, updated_at = $3, user_id = $4 WHERE id = $5", resume.S3Key, resume.CreatedAt, resume.UpdatedAt, resume.UserID, resume.ID)
	return err
}

func (r *ResumeRepository) Delete(id int) error {
	err := r.db.QueryRow("DELETE FROM resumes WHERE id = $1", id).Scan()
	return err
}
