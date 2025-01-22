-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS brewed_recipes (
    recipe_id UUID NOT NULL,
    witch_id UUID NOT NULL,
    status VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id),
    FOREIGN KEY (witch_id) REFERENCES witches(id)
    );
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS brewed_recipes;
-- +goose StatementEnd
