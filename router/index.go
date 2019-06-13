package router

import (
	"encoding/json"
	"net/http"
)

// IndexData is ...
type IndexData struct {
	Data string `json:"data"`
}

// Index is ...
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(IndexData{"Welcome to the Mealplanner app"})
	}
}
