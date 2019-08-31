package restaurants_owners

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

func RestaurantOwnerSchema(db *sql.DB) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    RestaurantOwnerQuery(db),
			Mutation: RestaurantOwnerMutation(db),
		})
}
