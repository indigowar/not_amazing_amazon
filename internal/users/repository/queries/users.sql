-- name: InsertUser :execresult
INSERT INTO users (
    id,
    passport, passport_salt,
    password, password_salt,
    displayed_name,
    phone_number, phone_number_salt
) VALUES (
    $1,
    $2, $3,
    $4, $5,
    $6,
    $7, $8
);
