package main

import (
	"fmt"
	"net/http"

	"github.com/toastbroad/mealplanner/database"
	"github.com/toastbroad/mealplanner/router"

	"github.com/go-pg/pg"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}

func main() {
	DB := database.Connect()
	DB.AddQueryHook(dbLogger{})

	err := database.CreateSchema()
	if err != nil {
		fmt.Println("Error ...", err)
	}

	r := router.GetRouter()

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8000", r)
}
