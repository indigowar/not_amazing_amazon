-- +goose Up
-- +goose StatementBegin

CREATE TABLE users(
    id UUID PRIMARY KEY,

    passport BYTEA NOT NULL,
    passport_salt CHAR(128) NOT NULL,
    UNIQUE(passport, passport_salt),

    password BYTEA NOT NULL,
    password_salt CHAR(128) NOT NULL,

    displayed_name VARCHAR(255) NOT NULL,

    phone_number BYTEA NOT NULL,
    phone_number_salt CHAR(128) NOT NULL,
    UNIQUE(phone_number, phone_number_salt),

    registration_date TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
