-- name: InsertUser :one
INSERT INTO users (email, password_hash)
VALUES ($1, $2)
RETURNING *;

-- name: ActivateUser :exec
UPDATE users
SET is_active = TRUE
WHERE user_id = $1;

-- name: FindByEmail :one
SELECT * FROM users
WHERE email = $1;