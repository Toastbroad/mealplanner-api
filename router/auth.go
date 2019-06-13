package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Auth is ...
func Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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
				fmt.Println("Token: " + token.Raw)

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					fmt.Println(claims["username"])
					fmt.Println(claims)
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		}
	}
}
