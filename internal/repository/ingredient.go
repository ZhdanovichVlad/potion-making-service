package repository

import (
	"context"
	"fmt"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
)

func (s *Storage) GetIngredients(ctx context.Context) ([]entity.Ingredient, error) {

	var ingredients []entity.Ingredient

	query := "SELECT * FROM ingredients"
	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, fmt.Errorf("preparation query error")
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("executing query error")
	}
	for rows.Next() {
		ingredient := entity.Ingredient{}
		err = rows.Scan(&ingredient)
		if err != nil {
			return nil, fmt.Errorf("scanning row error")
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}
