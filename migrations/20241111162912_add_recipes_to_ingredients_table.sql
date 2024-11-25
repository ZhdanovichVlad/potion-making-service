-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS recipes_to_ingredients (
    id_recipe BIGINT NOT NULL,
    id_ingredient BIGINT NOT NULL,
    amount int NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE IF EXISTS recipes_to_ingredients
-- +goose StatementEnd
