package categories

import (
	"../../db"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
)

func CategoryQueryType(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"categories":             getAllCategories(dataBase),
				"categoriesByRestaurant": getCategoriesByRestaurant(dataBase),
			},
		})
}

func getCategoriesByRestaurant(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(CategoryType),
		Args: graphql.FieldConfigArgument{
			"restaurant": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			restaurant, restaurantOk := p.Args["restaurant"].(int)

			if !restaurantOk {
				return nil, errors.New("restaurant not provided")
			}

			return db.GetCategoriesByRestaurant(dataBase, restaurant)
		},
		Description: "Get all categories by restaurant id",
	}
}

func getAllCategories(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(CategoryType),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return db.GetAllCategories(dataBase)
		},
		Description: "Get all categories",
	}
}
