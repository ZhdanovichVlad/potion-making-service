package repository

import (
	"context"
	"fmt"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
)

func (s *Storage) GetRecipes(ctx context.Context) ([]entity.Recipe, error) {

	var recipes []entity.Recipe
	query := "SELECT * FROM recipes"
	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, fmt.Errorf("preparation query error: %w", err)
	}

	rows, err := stmt.QueryContext(ctx)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("executing query error: %w", err)
	}
	for rows.Next() {
		recipe := entity.Recipe{}
		err = rows.Scan(&recipe.Id, &recipe.Name, &recipe.Description, &recipe.BrewTimeSeconds)
		if err != nil {
			return nil, fmt.Errorf("scanning row error: %w", err)
		}
		recipes = append(recipes, recipe)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return recipes, nil
}

func (s Storage) SaveRecipe(ctx context.Context, recipe entity.Recipe) error {
	query := "INSERT INTO recipes (name, description, brew_time_seconds) VALUES ($1, $2, $3)"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("preparation query error: %w", err)
	}

	_, err = stmt.ExecContext(ctx, recipe.Name, recipe.Description, recipe.BrewTimeSeconds)
	if err != nil {
		return fmt.Errorf("saving recipe error: %w", err)
	}
	return nil
}
