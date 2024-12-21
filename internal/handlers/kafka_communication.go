package handlers

import (
	"context"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
	"github.com/goccy/go-json"
)

func (a *PotionAPIServer) SaveRecipe(ctx context.Context, binary []byte) error {

	var recipe entity.Recipe

	err := json.Unmarshal(binary, &recipe)
	if err != nil {
		return err
	}

	if err = a.repo.SaveRecipe(ctx, recipe); err != nil {
		return err
	}

	return nil
}

func (a *PotionAPIServer) SaveIngredient(ctx context.Context, binary []byte) error {
	var ingredient entity.Ingredient
	err := json.Unmarshal(binary, &ingredient)
	if err != nil {
		return err
	}
	err = a.repo.SaveIngredient(ctx, ingredient)
	if err != nil {
		return err
	}
	return nil
}
