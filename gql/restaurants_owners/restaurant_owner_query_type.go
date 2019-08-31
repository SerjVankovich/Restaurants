package restaurants_owners

import (
	"../../db"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
)

func RestaurantOwnerQuery(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"owners": &graphql.Field{
					Type:        graphql.NewList(RestaurantOwnerType),
					Description: "Get all restaurant owners",
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return db.GetAllOwners(dataBase)
					},
				},
				"ownerById": &graphql.Field{
					Type:        RestaurantOwnerType,
					Description: "Get restaurant owner by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						id, idOk := p.Args["id"].(int)

						if !idOk {
							return nil, errors.New("ID not provided")
						}

						return db.GetOwnerById(dataBase, int32(id))
					},
				},
			},
		})
}
