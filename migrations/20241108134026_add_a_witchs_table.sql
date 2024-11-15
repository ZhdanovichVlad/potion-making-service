-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS witchs (
    id BIGSERIAL NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL
    );
CREATE INDEX IF NOT EXISTS ind_witchs_name ON witchs(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS ind_witchs_name
DROP TABLE IF EXISTS witchs
-- +goose StatementEnd
