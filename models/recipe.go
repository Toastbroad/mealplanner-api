package models

// Recipe is ...
type Recipe struct {
	ID            string        `json:"id"`
	Name          string        `json:"name" sql:",unique"`
	Source        string        `json:"source"`
	Category      string        `json:"category"`
	Author        string        `json:"author"`
	IngredientIDs []string      `json:"ingredientIDs,omitempty"`
	Ingredients   []*Ingredient `json:"ingredients,omitempty" pg:"many2many:recipe_to_ingredients"`
}
