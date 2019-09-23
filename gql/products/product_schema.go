package products

import (
	"database/sql"
	"github.com/graphql-go/graphql"
	"net/http"
)

func ProductSchema(dataBase *sql.DB, request *http.Request) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    ProductQueryType(dataBase, request),
			Mutation: ProductMutationType(dataBase, request),
		})
}
