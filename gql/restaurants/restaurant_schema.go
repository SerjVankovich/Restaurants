package restaurants

import (
	"database/sql"
	"github.com/graphql-go/graphql"
	"net/http"
)

func RestaurantSchema(dataBase *sql.DB, request *http.Request) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    RestaurantQuery(dataBase, request),
			Mutation: RestaurantMutation(dataBase, request),
		})
}
