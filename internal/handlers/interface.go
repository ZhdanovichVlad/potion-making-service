package handlers

import (
	"context"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
)

type Repository interface {
	GetIngredients(ctx context.Context) ([]entity.Ingredient, error)
	GetRecipes(ctx context.Context) ([]entity.Recipe, error)
}
