package processor

import (
	"context"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
)

type RecipesRepository interface {
	GetRecipes(ctx context.Context) ([]entity.Recipe, error)
}

type IngredientRepository interface {
	GetIngredients(ctx context.Context) ([]entity.Ingredient, error)
}
