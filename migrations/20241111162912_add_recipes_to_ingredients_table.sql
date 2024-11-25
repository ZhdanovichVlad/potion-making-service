-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS recipes_to_ingredients (
    recipe_id BIGINT NOT NULL,
    ingredient_id BIGINT NOT NULL,
    amount int NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE IF EXISTS recipes_to_ingredients;
-- +goose StatementEnd
