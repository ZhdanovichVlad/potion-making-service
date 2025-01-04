package repository

import (
	"context"
	"fmt"

	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"

	"github.com/jackc/pgx/v4"
)

const (
	getRecipes = "get_recipes"
	saveRecipe = "save_recipe"
)

type recipesStorage struct {
	db *pgx.Conn
}

func NewRecipesStorage(db *pgx.Conn) *recipesStorage {
	return &recipesStorage{db: db}
}

func (s *recipesStorage) Close(ctx context.Context) {
	s.db.Close(ctx)
}

func (s *recipesStorage) GetRecipes(ctx context.Context) ([]entity.Recipe, error) {

	var recipes []entity.Recipe
	query := "SELECT * FROM recipes"

	_, err := s.db.Prepare(ctx, getRecipes, query)
	if err != nil {
		return nil, fmt.Errorf("preparation query error: %w", err)
	}

	rows, err := s.db.Query(ctx, getRecipes)
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

func (s *recipesStorage) SaveRecipe(ctx context.Context, recipe entity.Recipe) error {
	query := "INSERT INTO recipes (name, description, brew_time_seconds) VALUES ($1, $2, $3)"

	_, err := s.db.Prepare(ctx, saveRecipe, query)
	if err != nil {
		return fmt.Errorf("preparation query error: %w", err)
	}

	_, err = s.db.Exec(ctx, saveRecipe, recipe.Name, recipe.Description, recipe.BrewTimeSeconds)
	if err != nil {
		return fmt.Errorf("saving recipe error: %w", err)
	}
	return nil
}
