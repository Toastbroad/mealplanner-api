package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	cookieutils "github.com/toastbroad/mealplanner/utils/cookie"
	jwtutils "github.com/toastbroad/mealplanner/utils/jwt"

	"github.com/dgrijalva/jwt-go"
)

// AuthErrorMsg is ...
type AuthErrorMsg struct {
	Message string `json:"errorMessage"`
}

// Authenticate ...
func Authenticate(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := r.Header.Get("Cookie")

		if strings.Contains(cookie, "=") {
			cookieValue := strings.Split(cookie, "=")[1]
			tokenString := strings.Trim(cookieValue, ";")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte("my_secret_key"), nil
			})

			if err == nil {
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // _ => claims
					var username string
					username = claims["username"].(string)
					tokenString, err := jwtutils.GetTokenString(username)
					if err != nil {
						w.WriteHeader(http.StatusUnauthorized)
						json.NewEncoder(w).Encode(AuthErrorMsg{"Not authenticated(a): " + err.Error()})
						return
					}
					expirationTime := time.Now().Add(5 * time.Minute)
					cookieutils.SetCookie(w, tokenString, expirationTime)
					next(w, r)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(AuthErrorMsg{"Not authenticated(a): " + err.Error()})
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(AuthErrorMsg{"Not authenticated(b): " + err.Error()})
			}

			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(AuthErrorMsg{"Not authenticated(b): No cookie"})
	}
}
