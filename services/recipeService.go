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
func CreateRecipe(parsedRecipe models.Recipe) (recipe models.Recipe, err error) {
	newRecipe := parsedRecipe
	newRecipe.ID = string(uuid.GenerateUUID())

	err = DB.Insert(&newRecipe)

	if err != nil {
		return newRecipe, err
	}

	for _, ingredientID := range parsedRecipe.IngredientIDs {
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
