-- name: InsertUser :exec
INSERT INTO users(id, passport, password, displayed_name, phone_number)
VALUES ($1, $2, $3, $4, $5);

-- name: SelectUserByID :one
SELECT * FROM users WHERE id = $1 AND is_deleted = FALSE;

-- name: SelectUserByPassport :one
SELECT * FROM users WHERE passport = $1 and is_deleted = FALSE;

-- name: SelectUserByPhoneNumber :one
SELECT * FROM users WHERE phone_number = $1 and is_deleted = FALSE;

-- name: DeleteUser :exec
UPDATE users
SET is_deleted = TRUE
WHERE id = $1;
