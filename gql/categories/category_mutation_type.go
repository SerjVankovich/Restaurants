package categories

import (
	"../../db"
	"../../models"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
)

func CategoryMutationType(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"addCategory": addCategory(dataBase),
			},
		})
}

func addCategory(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: CategoryType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"restaurant": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			name, nameOk := p.Args["name"].(string)
			restaurant, restaurantOk := p.Args["restaurant"].(int)

			if !nameOk {
				return nil, errors.New("name not provided")
			}

			if !restaurantOk {
				return nil, errors.New("restaurant not provided")
			}

			category := models.Category{Name: name, Restaurant: restaurant}

			err := db.AddCategory(dataBase, category)

			if err != nil {
				return nil, err
			}

			return &category, nil
		},
		Description: "Add one category",
	}
}
