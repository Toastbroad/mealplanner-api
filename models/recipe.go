package models

type Recipe struct {
	Id          []byte        `json:"id"`
	Name        string        `json:"name"`
	Source      string        `json:"source"`
	Author      string        `json:"author"`
	Ingredients []*Ingredient `json:"ingredients" pg:"many2many:recipe_to_ingredient"`
}
