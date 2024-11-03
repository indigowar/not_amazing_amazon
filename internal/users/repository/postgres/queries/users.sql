-- name: InsertUser :exec
INSERT INTO users(id, phone_number, password, displayed_name)
VALUES ($1, $2, $3, $4);

-- name: SelectUserByID :one
SELECT * FROM users WHERE id = $1 AND is_deleted = FALSE;

-- name: SelectUserByPhoneNumber :one
SELECT * FROM users WHERE phone_number = $1 and is_deleted = FALSE;

-- name: DeleteUser :exec
UPDATE users
SET is_deleted = TRUE
WHERE id = $1;
