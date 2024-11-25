-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS brewed_recipes (
    id_Recipe BIGINT NOT NULL,
    id_witch BIGINT NOT NULL,
    status VARCHAR NOT NULL,
    createdat TIMESTAMP NOT NULL,
    updatetat TIMESTAMP NOT NULL
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS brewed_recipes
-- +goose StatementEnd
