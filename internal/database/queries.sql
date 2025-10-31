-- ============================================
-- User Queries
-- ============================================

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetAllUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: CreateUser :one
INSERT INTO users (id, first_name, last_name, email, hashed_password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET first_name = $2,
    last_name = $3,
    email = $4,
    hashed_password = $5,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: VerifyUser :execrows
UPDATE users
SET verified = TRUE,
    updated_at = NOW()
WHERE id = $1;

-- ============================================
-- User Verification Link Queries
-- ============================================

-- name: CreateVerifyLink :exec
INSERT INTO user_verify_link (user_id, token, expires_at)
VALUES ($1, $2, $3);

-- name: DeleteVerifyLink :exec
DELETE FROM user_verify_link
WHERE user_id = $1;

-- name: GetVerifyLinkByToken :one
SELECT * FROM user_verify_link
WHERE token = $1 AND expires_at > NOW();

-- name: GetVerifyLinkByUserID :one
SELECT * FROM user_verify_link
WHERE user_id = $1;

-- ============================================
-- Resume Queries
-- ============================================

-- name: GetResume :one
SELECT * FROM resumes
WHERE id = $1;

-- name: GetAllResumes :many
SELECT * FROM resumes
ORDER BY created_at DESC;

-- name: GetResumesByUserID :many
SELECT * FROM resumes
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: CreateResume :one
INSERT INTO resumes (s3_key, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateResume :exec
UPDATE resumes
SET s3_key = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteResume :exec
DELETE FROM resumes
WHERE id = $1;

-- ============================================
-- Refresh Token Queries
-- ============================================

-- name: GetRefreshTokenByToken :one
SELECT * FROM refresh_token
WHERE token = $1;

-- name: GetRefreshTokenByUserID :one
SELECT * FROM refresh_token
WHERE user_id = $1;

-- name: CreateRefreshToken :one
INSERT INTO refresh_token (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_token
WHERE user_id = $1;

-- name: DeleteRefreshTokenByToken :exec
DELETE FROM refresh_token
WHERE token = $1;

-- name: DeleteExpiredRefreshTokens :exec
DELETE FROM refresh_token
WHERE expires_at < NOW();
