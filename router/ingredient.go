package router

import (
	"fmt"
	"net/http"

	"../services"
)

func Ingredient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		createIngredient(w, r)
		return
	}
}

func createIngredient(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	name := r.FormValue("name")
	ingredient, err := services.CreateIngredient(name)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error occured trying to create ingredient", 400)
		return
	}

	w.Header().Set("Location", "/ingredient/"+string(ingredient.Id[:]))
	w.WriteHeader(http.StatusCreated)
}
