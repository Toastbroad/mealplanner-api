package services

import "github.com/toastbroad/mealplanner/database"
import "github.com/toastbroad/mealplanner/models"

// CreateUser is ...
func CreateUser(name string, pw string) (user models.User, err error) {
	DB := database.Connect()

	newUser := models.User{
		ID:       "newid",
		UserName: name,
		Password: pw,
	}

	err = DB.Insert(&newUser)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
