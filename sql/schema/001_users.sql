-- to run a migration navigate to the folder containing the 
-- migration you want to run and 
-- run this command 
-- goose postgres <connection string> up

-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    "name" VARCHAR(255) UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;
