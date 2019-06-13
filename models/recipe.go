package models

// Recipe is ...
type Recipe struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Source      string        `json:"source"`
	Author      string        `json:"author"`
	Ingredients []*Ingredient `json:"ingredients,omitempty" pg:"many2many:recipe_to_ingredients"`
}
