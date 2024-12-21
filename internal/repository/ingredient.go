package repository

import (
	"context"
	"fmt"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
)

func (s *Storage) GetIngredients(ctx context.Context) ([]entity.Ingredient, error) {

	var ingredients []entity.Ingredient

	query := "SELECT * FROM ingredients"
	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, fmt.Errorf("query preparation error: %w", err)
	}

	rows, err := stmt.QueryContext(ctx)
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

func (s *Storage) SaveIngredient(ctx context.Context, ingredient entity.Ingredient) error {
	query := "UNSERT INTO ingredients (name, description) VALUES ($1, $2)"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("preparation query error: %w", err)
	}

	_, err = stmt.ExecContext(ctx, ingredient.Name, ingredient.Description)
	if err != nil {
		return fmt.Errorf("executing query error: %w", err)
	}

	return nil
}
