-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS recipes_to_ingredients (
    recipe_id UUID  REFERENCES recipes(id),
    ingredient_id UUID REFERENCES ingredients(id),
    amount INT NOT NULL,
    PRIMARY KEY (recipe_id, ingredient_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE TABLE IF EXISTS recipes_to_ingredients;
-- +goose StatementEnd
