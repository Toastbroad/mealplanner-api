package models

type Ingredient struct {
	Id      []byte    `json:"id"`
	Name    string    `json:"name"`
	Recipes []*Recipe `json:"recipes" pg:"many2many:recipe_to_ingredient"`
}
