#!/bin/bash

# This script is used for executing all migrations

TEMP_DIR=$(mktemp -d)

find ./ -path "*/migrations/*.sql" | sort | while IFS= read -r migration; do
    cp "$migration" "$TEMP_DIR/$(basename "$migration")"
done

goose -dir "$TEMP_DIR" postgres "host=127.0.0.1 port=5432 user=myuser password=mypassword database=mydatabase sslmode=disable" status

goose -dir "$TEMP_DIR" postgres "host=127.0.0.1 port=5432 user=myuser password=mypassword database=mydatabase sslmode=disable" up

rm -rf "$TEMP_DIR"

