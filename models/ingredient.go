package models

// Ingredient is ...
type Ingredient struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Recipes []*Recipe `json:"recipes,omitempty" pg:"many2many:recipe_to_ingredients"`
}
