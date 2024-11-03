-- +goose Up
-- +goose StatementBegin

CREATE TABLE users(
    id UUID PRIMARY KEY,
    phone_number VARCHAR(32) NOT NULL UNIQUE,
    password BYTEA NOT NULL UNIQUE,
    displayed_name VARCHAR(255) NOT NULL,
    registration_date TIMESTAMP NOT NULL DEFAULT now(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
