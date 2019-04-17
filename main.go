package main

import (
	"fmt"
	"net/http"

	"./database"
	"./router"
)

func main() {
	err := database.CreateSchema()
	if err != nil {
		fmt.Println("Error ...", err)
	}

	r := router.GetRouter()
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8000", r)
}
