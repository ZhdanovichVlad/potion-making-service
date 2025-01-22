package processor

import (
	"context"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
)

type RecipesRepository interface {
	GetRecipes(ctx context.Context) ([]entity.Recipe, error)
	SaveRecipeAndIngredient(ctx context.Context, newRecipe entity.CreateRecipe) error
}

type IngredientRepository interface {
	GetIngredients(ctx context.Context) ([]entity.Ingredient, error)
}
