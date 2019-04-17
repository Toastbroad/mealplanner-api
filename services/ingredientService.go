package services

import (
	"../database"
	"../models"
	"../utils/uuid"
)

func CreateIngredient(nameParam string) (err error) {
	DB := database.Connect()

	err = DB.Insert(&models.Ingredient{
		Id:   uuid.GenerateUUID(),
		Name: nameParam,
	})

	if err != nil {
		return err
	}

	return nil
}
