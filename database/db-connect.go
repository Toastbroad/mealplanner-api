package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"../models"
)

func Connect() *pg.DB {
	return DB
}

func CreateSchema() error {
	for _, model := range []interface{}{
		(*models.Ingredient)(nil),
		(*models.Recipe)(nil),
		(*models.RecipeToIngredient)(nil),
	} {
		err := DB.CreateTable(model, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
