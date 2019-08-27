package restaurants

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

func RestaurantSchema(dataBase *sql.DB) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    nil,
			Mutation: nil,
		})
}
