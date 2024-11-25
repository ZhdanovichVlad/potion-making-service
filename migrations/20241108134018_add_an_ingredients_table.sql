-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ingridients (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL
    );
CREATE INDEX IF NOT EXISTS ind_ingridients_name ON ingridients (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS ind_ingridients_name;
DROP TABLE IF EXISTS ingridients;
-- +goose StatementEnd
