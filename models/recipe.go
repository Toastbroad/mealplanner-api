package models

type Recipe struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Source      string        `json:"source"`
	Author      string        `json:"author"`
	Ingredients []*Ingredient `json:"ingredients" pg:"many2many:recipe_to_ingredients"`
}
