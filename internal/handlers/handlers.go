package handlers

import (
	"context"
	"github.com/ZhdanovichVlad/potion-making-service/branches/generated/openapi"
	"net/http"
)

type PotionAPIServer struct {
	openapi.DefaultAPIController
}

// NewDefaultAPIService creates a default api service
func NewPotionAPIServer() *PotionAPIServer {
	return &PotionAPIServer{}
}

// ImplResponse
func (s *PotionAPIServer) GetAllRecipes(ctx context.Context) (openapi.ImplResponse, error) {

	// TODO - update GetAllRecipes with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []ApiResponseRecipes{}) or use other options such as http.Ok ...
	// return Response(200, []ApiResponseRecipes{}), nil

	//ApiResponseIngredients
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	// TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	// return Response(500, nil),nil

	return openapi.Response(http.StatusOK, "all ok"), nil
}

// GetAllIngredients - returns an array of ingredients
func (s *PotionAPIServer) GetAllIngredients(ctx context.Context) (openapi.ImplResponse, error) {

	// TODO - update GetAllIngredients with the required logic for this service method.

	// TODO: Uncomment the next line to return response Response(200, []ApiResponseIngredients{}) or use other options such as http.Ok ...
	// return Response(200, []ApiResponseIngredients{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	// TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	// return Response(500, nil),nil

	return openapi.Response(http.StatusOK, "all ok"), nil
}
