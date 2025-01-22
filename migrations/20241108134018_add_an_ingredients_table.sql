-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ingredients (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) UNIQUE NOT NULL,
    description VARCHAR(500) NOT NULL
    );
CREATE INDEX IF NOT EXISTS ind_ingridients_name ON ingredients(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS ind_ingridients_name;
DROP TABLE IF EXISTS ingredients;
-- +goose StatementEnd
