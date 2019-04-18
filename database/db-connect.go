package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"github.com/toastbroad/mealplanner/models"
)

func Connect() *pg.DB {
	return DB
}

func CreateSchema() error {
	orm.RegisterTable((*models.RecipeToIngredient)(nil))

	for _, model := range []interface{}{
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
