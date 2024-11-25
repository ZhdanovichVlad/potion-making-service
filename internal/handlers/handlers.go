package handlers

import (
	"context"
	"fmt"
	"github.com/ZhdanovichVlad/potion-making-service/branches/generated/openapi"
	"net/http"
)

type PotionAPIServer struct {
	repo Repository
	openapi.DefaultAPIController
}

// NewDefaultAPIService creates a default api service
func NewPotionAPIServer(repo Repository) *PotionAPIServer {
	return &PotionAPIServer{repo: repo}
}

// ImplResponse
func (s *PotionAPIServer) GetAllRecipes(ctx context.Context) (openapi.ImplResponse, error) {

	// TODO - update GetAllRecipes with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []ApiResponseRecipes{}) or use other options such as http.Ok ...
	// return Response(200, []ApiResponseRecipes{}), nil

	recipes, err := s.repo.GetRecipes(ctx)
	if err != nil {
		return openapi.Response(500, nil), fmt.Errorf("failed to get recipes: %w", err)
	}

	apiResponse := make([]openapi.Recipe, 0, len(recipes))
	for _, recipe := range recipes {
		apiResponse = append(apiResponse, openapi.Recipe{recipe.Id, recipe.Name, recipe.Description, recipe.BrewTimeSeconds})
	}

	//ApiResponseIngredients
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return openapi.Response(http.StatusOK, apiResponse), err
}

// GetAllIngredients - returns an array of ingredients
func (s *PotionAPIServer) GetAllIngredients(ctx context.Context) (openapi.ImplResponse, error) {

	// TODO - update GetAllIngredients with the required logic for this service method.

	Ingredients, err := s.repo.GetIngredients(ctx)
	if err != nil {
		return openapi.Response(500, nil), fmt.Errorf("failed to get Ingredients: %w", err)
	}

	apiResponse := make([]openapi.Ingredient, 0, len(Ingredients))
	for _, ingredient := range Ingredients {
		apiResponse = append(apiResponse, openapi.Ingredient{ingredient.Id, ingredient.Name, ingredient.Description})
	}

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return openapi.Response(http.StatusOK, apiResponse), nil
}
