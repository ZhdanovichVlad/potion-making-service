-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS brewed_recipes (
    recipe_id BIGINT NOT NULL,
    witch_id BIGINT NOT NULL,
    status VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS brewed_recipes;
-- +goose StatementEnd
