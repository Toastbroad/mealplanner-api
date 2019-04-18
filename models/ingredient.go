package models

type Ingredient struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Recipes []*Recipe `json:"recipes" pg:"many2many:recipe_to_ingredients"`
}
