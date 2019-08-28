package categories

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

func CategorySchema(dataBase *sql.DB) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    CategoryQueryType(dataBase),
			Mutation: CategoryMutationType(dataBase),
		})
}
