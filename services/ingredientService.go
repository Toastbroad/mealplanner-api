package services

import (
	"github.com/toastbroad/mealplanner/database"
	"github.com/toastbroad/mealplanner/models"
	"github.com/toastbroad/mealplanner/utils/uuid"
)

func CreateIngredient(nameParam string) (ingredient models.Ingredient, err error) {
	DB := database.Connect()
	newIngredient := models.Ingredient{
		Id:   string(uuid.GenerateUUID()),
		Name: nameParam,
	}

	err = DB.Insert(&newIngredient)

	if err != nil {
		return newIngredient, err
	}

	return newIngredient, nil
}
