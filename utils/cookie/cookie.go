package cookieutils

import (
	"net/http"
	"time"
)

// SetCookie is ...
func SetCookie(w http.ResponseWriter, tokenString string, expirationTime time.Time) {
	//expirationTime := time.Now().Add(5 * time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
