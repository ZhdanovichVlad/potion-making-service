package consumer

import "context"

type RecipesSaver interface {
	SaveRecipe(ctx context.Context, binary []byte) error
}

