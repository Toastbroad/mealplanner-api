package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRouter() http.Handler {
	var jsonRouterConfig = map[string]func(w http.ResponseWriter, r *http.Request){
		"/":            Index,
		"/ingredient":  Ingredient,
		"/recipe":      Recipe,
		"/recipe/{id}": RecipeById,
	}

	r := mux.NewRouter().StrictSlash(true)

	for path, handleFunc := range jsonRouterConfig {
		path := path
		handleFunc := handleFunc

		r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			handleFunc(w, r)
		})
	}

	return r
}
