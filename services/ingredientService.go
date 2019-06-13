package services

import (
	"github.com/toastbroad/mealplanner/database"
	"github.com/toastbroad/mealplanner/models"
	"github.com/toastbroad/mealplanner/utils/uuid"
)

// CreateIngredient is ...
func CreateIngredient(nameParam string) (ingredient models.Ingredient, err error) {
	DB := database.Connect()
	newIngredient := models.Ingredient{
		ID:   string(uuid.GenerateUUID()),
		Name: nameParam,
	}

	err = DB.Insert(&newIngredient)

	if err != nil {
		return newIngredient, err
	}

	return newIngredient, nil
}
