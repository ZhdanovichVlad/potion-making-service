package handlers

import "github.com/ZhdanovichVlad/potion-making-service/branches/generated/openapi"

type PotionAPIServer struct {
	repo Repository
	openapi.DefaultAPIController
}

// NewDefaultAPIService creates a default api service
func NewPotionAPIServer(repo Repository) *PotionAPIServer {
	return &PotionAPIServer{repo: repo}
}
