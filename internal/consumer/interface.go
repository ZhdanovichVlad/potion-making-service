package consumer

import (
	"context"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
)

type RecipesSaver interface {
	SaveRecipe(ctx context.Context, recipe entity.CreateRecipe) error
}
