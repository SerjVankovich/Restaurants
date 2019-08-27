package restaurants

import (
	"../../db"
	"database/sql"
	"github.com/graphql-go/graphql"
)

func RestaurantQuery(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"restaurants": getAllRestaurants(dataBase),
			},
		},
	)
}

func getAllRestaurants(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(RestaurantType),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return db.GetAllRestaurants(dataBase)
		},
		Description: "Get all restaurants",
	}
}
