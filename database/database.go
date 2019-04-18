package database

import "github.com/go-pg/pg"

var DB = pg.Connect(&pg.Options{
	User:     "mealplanner",
	Password: "mealplanner",
})
