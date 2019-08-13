package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/toastbroad/mealplanner-api/models"
	"github.com/toastbroad/mealplanner-api/services"
)

// Ingredient is ...
func Ingredient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		createIngredient(w, r)
		return
	}

	if r.Method == http.MethodGet {
		getIngredients(w, r)
	}
}

func getIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := services.GetIngredients()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error occured trying to create ingredient", 400)
		return
	}
	fmt.Println("ingredients: ", ingredients)
	json.NewEncoder(w).Encode(ingredients)
}

func createIngredient(w http.ResponseWriter, r *http.Request) {
	ingredient := &models.Ingredient{}
	json.NewDecoder(r.Body).Decode(ingredient)
	newIngredient, err := services.CreateIngredient(*ingredient)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error occured trying to create ingredient", 400)
		return
	}

	w.Header().Set("Location", "/ingredient/"+string(newIngredient.ID[:]))
	w.WriteHeader(http.StatusCreated)
}
