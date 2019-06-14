package models

// User is ...
type User struct {
	ID       string `json:"id"`
	UserName string `json:"username" sql:",unique"`
	Password string `json:"password"`
}
