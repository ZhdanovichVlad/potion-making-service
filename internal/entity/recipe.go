package entity

type Recipe struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	BrewTimeSeconds int64  `json:"brewTimeSeconds"`
}

type CreateRecipe struct {
	Ingredients []Ingredient
	Recipe
}
