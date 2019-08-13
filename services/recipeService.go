package services

import (
	"github.com/toastbroad/mealplanner-api/database"
	"github.com/toastbroad/mealplanner-api/models"
	"github.com/toastbroad/mealplanner-api/utils/uuid"
)

// DB is ...
var DB = database.Connect()

// GetRecipes is ...
func GetRecipes() (recipes []models.Recipe, err error) {
	err = DB.Model(&recipes).Relation("Ingredients").Select()

	if err != nil {
		return recipes, err
	}

	return recipes, err
}

// GetRecipeByID is ...
func GetRecipeByID(id string) (recipe models.Recipe, err error) {
	err = DB.Model(&recipe).Relation("Ingredients").Where(`id='` + id + `'`).Select()

	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

// CreateRecipe is ...
func CreateRecipe() (recipe models.Recipe, err error) {
	ingredientIDs := []string{
		"30299911-6EF3-468B-8008-8C25F4247610",
		"8DF46963-046E-4ECD-B9AE-82EF5F50C2B2",
	}

	newRecipe := models.Recipe{
		ID:     string(uuid.GenerateUUID()),
		Name:   "Super awesome recipe II",
		Source: "Delightful Cookbook Vol II",
		Author: "Oliver Broad",
	}

	err = DB.Insert(&newRecipe)

	if err != nil {
		return newRecipe, err
	}

	for _, ingredientID := range ingredientIDs {
		err := DB.Insert(&models.RecipeToIngredient{
			RecipeID:     newRecipe.ID,
			IngredientID: ingredientID,
		})

		if err != nil {
			return newRecipe, err
		}
	}

	return newRecipe, nil
}
