package router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/toastbroad/mealplanner-api/services"
	cookieutils "github.com/toastbroad/mealplanner-api/utils/cookie"
	jwtutils "github.com/toastbroad/mealplanner-api/utils/jwt"

	"golang.org/x/crypto/bcrypt"
)

// Credentials is a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login is ...
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := services.GetUserByName(creds.Username)

	expectedPassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(expectedPassword), []byte(creds.Password))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, err := jwtutils.GetTokenString(creds.Username)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	cookieutils.SetCookie(w, tokenString, expirationTime)
}
