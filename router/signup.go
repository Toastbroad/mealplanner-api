package router

import (
	"encoding/json"
	"net/http"

	"github.com/toastbroad/mealplanner/services"
	errorutils "github.com/toastbroad/mealplanner/utils/error"

	"golang.org/x/crypto/bcrypt"
)

// Signup is ...
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorutils.ErrorMsg{Message: "Http method not allowed."})
		return
	}
	// Parse and decode the request body into a new `Credentials` instance
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorutils.ErrorMsg{Message: "Creating user failed. Parsing request body unsuccessful: " + err.Error()})
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	// Next, insert the username, along with the hashed password into the database
	newUser, err := services.CreateUser(creds.Username, string(hashedPassword))
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorutils.ErrorMsg{Message: "Creating user failed: " + err.Error()})
		return
	}

	// todo: set location header for created user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
