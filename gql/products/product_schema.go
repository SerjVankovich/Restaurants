package products

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

func ProductSchema(dataBase *sql.DB) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    ProductQueryType(dataBase),
			Mutation: ProductMutationType(dataBase),
		})
}
