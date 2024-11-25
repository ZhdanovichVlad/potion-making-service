-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS recipes (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    brew_time_seconds int
);
CREATE INDEX IF NOT EXISTS ind_recipes_name ON recipes (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS ind_recipes_name;
DROP TABLE IF EXISTS recipes;
-- +goose StatementEnd
