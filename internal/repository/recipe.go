package repository

import (
	"context"
	"fmt"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/entity"
)

func (s *Storage) GetRecipes(ctx context.Context) ([]entity.Recipe, error) {

	var recipes []entity.Recipe
	query := "SELECT * FROM recipes"
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
		recipe := entity.Recipe{}
		err = rows.Scan(&recipe)
		if err != nil {
			return nil, fmt.Errorf("scanning row error")
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}
