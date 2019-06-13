package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"github.com/toastbroad/mealplanner/models"
)

// Connect is ...
func Connect() *pg.DB {
	return DB
}

// CreateSchema is ...
func CreateSchema() error {
	orm.RegisterTable((*models.RecipeToIngredient)(nil))

	for _, model := range []interface{}{
		(*models.User)(nil),
		(*models.RecipeToIngredient)(nil),
		(*models.Recipe)(nil),
		(*models.Ingredient)(nil),
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
