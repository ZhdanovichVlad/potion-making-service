package processor

import (
	"context"
	"fmt"
	"github.com/ZhdanovichVlad/potion-making-service/branches/generated/openapi"
)

type IngredientAPIServer struct {
	repo IngredientRepository
	openapi.IngredientAPIService
}

// IngredientAPIServer creates a Ingredient api service
func NewIngredientAPIServer(repo IngredientRepository) *IngredientAPIServer {
	return &IngredientAPIServer{repo: repo}
}

// GetAllIngredients - returns an array of ingredients
func (s *IngredientAPIServer) GetAllIngredients(ctx context.Context) (openapi.ImplResponse, error) {

	ingredients, err := s.repo.GetIngredients(ctx)
	if err != nil {
		return openapi.Response(500, nil), fmt.Errorf("failed to get Ingredients: %w", err)
	}

	apiResponse := make([]openapi.Ingredient, 0, len(ingredients))
	for _, ingredient := range ingredients {
		apiResponse = append(apiResponse, openapi.Ingredient{ingredient.Id, ingredient.Name, ingredient.Description})
	}

	return openapi.Response(200, apiResponse), nil
}
