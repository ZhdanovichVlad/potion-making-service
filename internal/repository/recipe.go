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

func (s *recipesStorage) SaveRecipeAndIngredient(ctx context.Context, newRecipe entity.CreateRecipe) error {
	queryInsertRecipe := "INSERT INTO recipes (name, description, brew_time_seconds) VALUES ($1, $2, $3)"
	queryInsertIngredient := "INSERT INTO ingredients (name, description) VALUES ($1, $2)"

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("transaction begin error")
	}

	_, err = tx.Prepare(ctx, saveRecipe, queryInsertRecipe)
	if err != nil {
		errRollback := tx.Rollback(ctx)
		if errRollback != nil {
			return fmt.Errorf(" rollback error: %w with an error when preparation query insert recipe error: %w", errRollback, err)
		}
		return fmt.Errorf("preparation query insert recipe error: %w", err)
	}

	_, err = tx.Exec(ctx, saveRecipe, newRecipe.Name, newRecipe.Description, newRecipe.BrewTimeSeconds)
	if err != nil {
		errRollback := tx.Rollback(ctx)
		if errRollback != nil {
			return fmt.Errorf(" rollback error: %w with an error when inserting the recipe.: %w", errRollback, err)
		}

		return fmt.Errorf("saving recipe error: %w", err)
	}

	_, err = tx.Prepare(ctx, saveIngredient, queryInsertIngredient)
	if err != nil {
		errRollback := tx.Rollback(ctx)
		if errRollback != nil {
			return fmt.Errorf(" rollback error: %w with an error when preparation query insert ingredient error: %w", errRollback, err)
		}
		return fmt.Errorf("preparation query nsert ingredient error: %w", err)
	}

	for _, ingredient := range newRecipe.Ingredients {

		_, err = tx.Exec(ctx, saveIngredient, ingredient.Name, ingredient.Description)
		if err != nil {
			errRollback := tx.Rollback(ctx)
			if errRollback != nil {
				return fmt.Errorf(" rollback error: %w with an error when inserting the ingredient.: %w", errRollback, err)
			}
			return fmt.Errorf("executing query error: %w", err)
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		errRollback := tx.Rollback(ctx)
		if errRollback != nil {
			return fmt.Errorf(" rollback error: %w with an error when commiting changes: %w", errRollback, err)
		}
		return fmt.Errorf("commiting changes error: %w", err)
	}

	return nil
}
