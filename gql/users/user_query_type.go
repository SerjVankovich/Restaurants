package users

import (
	"../../db"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
)

func UserQuery(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"users": &graphql.Field{
					Type:        graphql.NewList(UserType),
					Description: "Get all users",
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return db.GetAllUsers(dataBase)
					},
				},
				"userById": &graphql.Field{
					Type:        UserType,
					Description: "Get user by ID",
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

						return db.GetUserById(dataBase, int32(id))
					},
				},
			},
		})
}
