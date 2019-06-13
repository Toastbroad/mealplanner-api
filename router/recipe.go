package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/toastbroad/mealplanner/services"
)

// Recipe is ...
func Recipe(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getRecipes(w, r)
		return
	}

	if r.Method == http.MethodPost {
		createRecipe(w, r)
		return
	}
}

// RecipeByID is ...
func RecipeByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	recipe, err := services.GetRecipeByID(id)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error occured trying to get recipes", 400)
		return
	}

	json.NewEncoder(w).Encode(recipe)
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := services.GetRecipes()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error occured trying to get recipes", 400)
		return
	}

	json.NewEncoder(w).Encode(recipes)
}

func createRecipe(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	recipe, err := services.CreateRecipe()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error occured trying to create recipe", 400)
		return
	}

	w.Header().Set("Location", "/recipe/"+string(recipe.ID[:]))
	w.WriteHeader(http.StatusCreated)
}
