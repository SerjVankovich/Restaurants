package orders

import (
	"database/sql"
	"github.com/graphql-go/graphql"
	"net/http"
)

func OrderSchema(dataBase *sql.DB, request *http.Request) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    OrderQueryType(dataBase, request),
			Mutation: OrderMutationType(dataBase, request),
		})
}
