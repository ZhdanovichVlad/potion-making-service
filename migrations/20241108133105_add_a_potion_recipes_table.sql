-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS recipes (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(36) NOT NULL,
    description VARCHAR NOT NULL,
    brew_time_seconds INT
);
CREATE INDEX IF NOT EXISTS ind_recipes_name ON recipes (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS ind_recipes_name;
DROP TABLE IF EXISTS recipes;
-- +goose StatementEnd
