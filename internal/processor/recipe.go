package processor

import (
	"Context"
	"fmt"
	"github.com/ZhdanovichVlad/potion-making-service/branches/generated/openapi"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
)

type RecipesAPIServer struct {
	repo RecipesRepository
	openapi.RecipeAPIController
}

// IngredientAPIServer creates a Ingredient api service
func NewRecipesAPIServer(repo RecipesRepository) *RecipesAPIServer {
	return &RecipesAPIServer{repo: repo}
}

// GetAllRecipes returns an array of Recipes
func (s *RecipesAPIServer) GetAllRecipes(ctx context.Context) (openapi.ImplResponse, error) {

	recipes, err := s.repo.GetRecipes(ctx)
	if err != nil {
		return openapi.ImplResponse{}, fmt.Errorf("failed to get recipes: %w", err)
	}

	apiResponse := make([]openapi.Recipe, 0, len(recipes))
	for _, recipe := range recipes {
		apiResponse = append(apiResponse, openapi.Recipe{recipe.Id, recipe.Name, recipe.Description, recipe.BrewTimeSeconds})
	}

	return openapi.Response(200, apiResponse), err
}

// Save the recipe and ingredients in the database
func (s *RecipesAPIServer) SaveRecipe(ctx context.Context, newRecipe entity.CreateRecipe) error {

	err := s.repo.SaveRecipeAndIngredient(ctx, newRecipe)
	if err != nil {
		return err
	}
	return nil
}
