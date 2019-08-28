package restaurants

import (
	"../../db"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
)

func RestaurantQuery(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"restaurants":       getAllRestaurants(dataBase),
				"restaurantById":    getRestaurantByID(dataBase),
				"restaurantsByName": getRestaurantsByName(dataBase),
			},
		},
	)
}

func getRestaurantsByName(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(RestaurantType),
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			name, nameOk := p.Args["name"].(string)
			if !nameOk {
				return nil, errors.New("name not provided")
			}

			return db.GetRestaurantsByName(dataBase, name)
		},
		Description: "Get restaurant by name",
	}
}

func getRestaurantByID(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: RestaurantType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id, idOk := p.Args["id"].(int)
			if !idOk {
				return nil, errors.New("id not provided")
			}
			restaurant, err := db.GetRestaurantById(dataBase, id)

			return restaurant, err
		},
		Description: "Get restaurant by id",
	}
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
