package services

import (
	"../database"
	"../models"
	"../utils/uuid"
)

func CreateRecipe() (err error) {
	DB := database.Connect()

	err = DB.Insert(&models.Recipe{
		Id:     uuid.GenerateUUID(),
		Name:   "Super awesome recipe",
		Source: "Delightful Cookbook Vol II",
		Author: "Oliver Broad",
	})

	if err != nil {
		return err
	}

	return nil
}
