package models

type CocktailIngredientModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
