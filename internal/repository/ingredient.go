package repository

import (
	"context"
	"fmt"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
	"github.com/jackc/pgx/v4"
)

const (
	getIngredients = "get_ingredients"
	saveIngredient = "save_ingredient"
)

type ingredientsStorage struct {
	db *pgx.Conn
}

func NewIngredientsStorage(db *pgx.Conn) *ingredientsStorage {
	return &ingredientsStorage{db: db}
}

func (s *ingredientsStorage) Close(ctx context.Context) {
	s.db.Close(ctx)
}

func (s *ingredientsStorage) GetIngredients(ctx context.Context) ([]entity.Ingredient, error) {

	var ingredients []entity.Ingredient

	query := "SELECT * FROM ingredients"               // Можно ли вынести в отдельный метод и инициализировать только 1 раз
	_, err := s.db.Prepare(ctx, getIngredients, query) // Можно ли вынести в отдельный метод и инициализировать только 1 раз
	if err != nil {
		return nil, fmt.Errorf("query preparation error: %w", err)
	} // Можно ли вынести в отдельный метод и инициализировать только 1 раз

	rows, err := s.db.Query(ctx, getIngredients)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("executing query error: %w", err)
	}
	for rows.Next() {
		ingredient := entity.Ingredient{}
		err = rows.Scan(&ingredient.Id, &ingredient.Name, &ingredient.Description)
		if err != nil {
			return nil, fmt.Errorf("scanning row error: %w", err)
		}
		ingredients = append(ingredients, ingredient)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return ingredients, nil
}

func (s *ingredientsStorage) SaveIngredient(ctx context.Context, ingredient entity.Ingredient) error {

	query := "INSERT INTO ingredients (name, description) VALUES ($1, $2)"

	_, err := s.db.Prepare(ctx, saveIngredient, query)
	if err != nil {
		return fmt.Errorf("preparation query error: %w", err)
	}

	_, err = s.db.Exec(ctx, saveIngredient, ingredient.Name, ingredient.Description)
	if err != nil {
		return fmt.Errorf("executing query error: %w", err)
	}

	return nil
}
