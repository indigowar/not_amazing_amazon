-- +goose Up
-- +goose StatementBegin

CREATE TABLE users(
    id UUID PRIMARY KEY,
    passport BYTEA NOT NULL UNIQUE,
    password BYTEA NOT NULL UNIQUE,
    displayed_name VARCHAR(255) NOT NULL,
    phone_number BYTEA NOT NULL UNIQUE,
    registration_date TIMESTAMP NOT NULL DEFAULT now(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
