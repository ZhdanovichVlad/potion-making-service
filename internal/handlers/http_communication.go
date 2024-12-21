package handlers

import (
	"context"
	"fmt"
	"github.com/ZhdanovichVlad/potion-making-service/branches/generated/openapi"
	"net/http"
)

// ImplResponse
func (s *PotionAPIServer) GetAllRecipes(ctx context.Context) (openapi.ImplResponse, error) {

	recipes, err := s.repo.GetRecipes(ctx)
	if err != nil {
		return openapi.Response(500, nil), fmt.Errorf("failed to get recipes: %w", err)
	}

	apiResponse := make([]openapi.Recipe, 0, len(recipes))
	for _, recipe := range recipes {
		apiResponse = append(apiResponse, openapi.Recipe{recipe.Id, recipe.Name, recipe.Description, recipe.BrewTimeSeconds})
	}

	return openapi.Response(http.StatusOK, apiResponse), err
}

// GetAllIngredients - returns an array of ingredients
func (s *PotionAPIServer) GetAllIngredients(ctx context.Context) (openapi.ImplResponse, error) {

	ingredients, err := s.repo.GetIngredients(ctx)
	if err != nil {
		return openapi.Response(500, nil), fmt.Errorf("failed to get Ingredients: %w", err)
	}

	apiResponse := make([]openapi.Ingredient, 0, len(ingredients))
	for _, ingredient := range ingredients {
		apiResponse = append(apiResponse, openapi.Ingredient{ingredient.Id, ingredient.Name, ingredient.Description})
	}

	return openapi.Response(http.StatusOK, apiResponse), nil
}
