package services

import (
	"github.com/toastbroad/mealplanner-api/database"
	"github.com/toastbroad/mealplanner-api/models"
	"github.com/toastbroad/mealplanner-api/utils/uuid"
)

func GetIngredients() (ingredients []models.Ingredient, err error) {
	DB := database.Connect()
	err = DB.Model(&ingredients).Select()
	return ingredients, err
}

// CreateIngredient is ...
func CreateIngredient(parsedIngredient models.Ingredient) (ingredient models.Ingredient, err error) {
	DB := database.Connect()
	newIngredient := models.Ingredient{
		ID:   string(uuid.GenerateUUID()),
		Name: parsedIngredient.Name,
	}

	err = DB.Insert(&newIngredient)

	if err != nil {
		return newIngredient, err
	}

	return newIngredient, nil
}
