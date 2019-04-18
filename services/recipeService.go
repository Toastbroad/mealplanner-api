package services

import (
	"../database"
	"../models"
	"../utils/uuid"
)

var DB = database.Connect()

func GetRecipes() (recipes []models.Recipe, err error) {
	_, err = DB.Query(&recipes, `SELECT * FROM recipes`)

	if err != nil {
		return recipes, err
	}

	return recipes, err
}

func CreateRecipe() (recipe models.Recipe, err error) {
	ingredients := []*models.Ingredient{
		&models.Ingredient{Id: "69509427-891D-4A6C-9BAF-77ACA62DBB15"},
	}
	//DB := database.Connect()
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
