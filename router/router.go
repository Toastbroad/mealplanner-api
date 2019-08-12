package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/toastbroad/mealplanner-api/middleware"
)

// GetRouter is ...
func GetRouter() http.Handler {
	var auth = middleware.Authenticate
	var jsonRouterConfig = map[string]func(w http.ResponseWriter, r *http.Request){
		"/":            Index,
		"/signup":      Signup,
		"/login":       Login,
		"/ingredient":  Ingredient,
		"/recipe":      Recipe,
		"/recipe/{id}": RecipeByID,
		"/test":        auth(Index),
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
