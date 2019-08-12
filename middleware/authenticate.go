package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	cookieutils "github.com/toastbroad/mealplanner-api/utils/cookie"
	errorutils "github.com/toastbroad/mealplanner-api/utils/error"
	jwtutils "github.com/toastbroad/mealplanner-api/utils/jwt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

// Authenticate ...
func Authenticate(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenString := cookie.Value

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
					json.NewEncoder(w).Encode(errorutils.ErrorMsg{Message: "Not authenticated: " + err.Error()})
					return
				}

				if time.Unix(int64(claims["exp"].(float64)), 0).Sub(time.Now()) <= 30*time.Second {
					fmt.Println("refreshing cookie")
					expirationTime := time.Now().Add(5 * time.Minute)
					cookieutils.SetCookie(w, tokenString, expirationTime)
				}
				context.Set(r, "username", username) // would make more sense to have user ID than username!?!?
				next(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorutils.ErrorMsg{Message: "Not authenticated: " + err.Error()})
		}
	}
}
