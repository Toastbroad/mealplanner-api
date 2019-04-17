package router

import (
	"fmt"
	"net/http"

	"../services"
)

func Recipe(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		createRecipe(w, r)
		return
	}
}

func createRecipe(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	err := services.CreateRecipe()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error occured trying to create ingredient", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
