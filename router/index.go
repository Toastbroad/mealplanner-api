package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
)

// IndexData is ...
type IndexData struct {
	Data string `json:"data"`
}

// Index is ...
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		username, ok := context.GetOk(r, "username")
		if ok {
			json.NewEncoder(w).Encode(IndexData{"Welcome to the Mealplanner app " + username.(string)})
			return
		}

		json.NewEncoder(w).Encode(IndexData{"Welcome to the Mealplanner app"})
	}
}
