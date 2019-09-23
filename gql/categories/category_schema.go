package categories

import (
	"database/sql"
	"github.com/graphql-go/graphql"
	"net/http"
)

func CategorySchema(dataBase *sql.DB, request *http.Request) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    CategoryQueryType(dataBase),
			Mutation: CategoryMutationType(dataBase, request),
		})
}
