package services

import (
	"../database"
	"../models"
	"../utils/uuid"
)

var DB = database.Connect()

func GetRecipes() (recipes []models.Recipe, err error) {
	err = DB.Model(&recipes).Column("id", "name", "author", "source").Select()

	if err != nil {
		return recipes, err
	}

	return recipes, err
}

func GetRecipeById(id string) (recipe models.Recipe, err error) {
	err = DB.Model(&recipe).Relation("Ingredients").Where(`id='` + id + `'`).Select()

	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

func CreateRecipe() (recipe models.Recipe, err error) {
	ingredients := []*models.Ingredient{
		&models.Ingredient{Id: "6DC6F2FF-406A-4A68-9B9C-05DC9E1D8017"},
		&models.Ingredient{Id: "A85E3914-41B0-4CC3-82C3-29D7099EFEAD"},
	}

	newRecipe := models.Recipe{
		Id:          string(uuid.GenerateUUID()),
		Name:        "Super awesome recipe",
		Source:      "Delightful Cookbook Vol II",
		Author:      "Oliver Broad",
		Ingredients: ingredients,
	}

	err = DB.Insert(&newRecipe)

	if err != nil {
		return newRecipe, err
	}

	for _, ingredient := range newRecipe.Ingredients {
		err := DB.Insert(&models.RecipeToIngredient{
			RecipeID:     newRecipe.Id,
			IngredientID: ingredient.Id,
		})

		if err != nil {
			return newRecipe, err
		}
	}

	return newRecipe, nil
}
