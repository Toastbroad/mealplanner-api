package services

import (
	"github.com/toastbroad/mealplanner-api/database"
	"github.com/toastbroad/mealplanner-api/models"
	"github.com/toastbroad/mealplanner-api/utils/uuid"
)

// CreateUser is ...
func CreateUser(name string, pw string) (user models.User, err error) {
	DB := database.Connect()

	newUser := models.User{
		ID:       uuid.GenerateUUID(),
		UserName: name,
		Password: pw,
	}

	err = DB.Insert(&newUser)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
