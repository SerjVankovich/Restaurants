package orders

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

func OrderSchema(dataBase *sql.DB) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    OrderQueryType(dataBase),
			Mutation: OrderMutationType(dataBase),
		})
}
