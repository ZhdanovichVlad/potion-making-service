-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS witches (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL
    );
CREATE INDEX IF NOT EXISTS ind_witches_name ON witches(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS ind_witches_name;
DROP TABLE IF EXISTS witches;
-- +goose StatementEnd
