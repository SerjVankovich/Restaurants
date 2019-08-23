package users

import (
	"database/sql"
	"github.com/graphql-go/graphql"
)

func UserSchema(db *sql.DB) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    UserQuery(db),
			Mutation: UserMutation(db),
		})
}
