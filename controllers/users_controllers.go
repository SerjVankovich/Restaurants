package controllers

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"net/http"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(
		graphql.Params{
			Schema:        schema,
			RequestString: query,
		},
	)

	return result
}

func GQLHandler(schema graphql.Schema) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		result := executeQuery(request.URL.Query().Get("query"), schema)
		_ = json.NewEncoder(writer).Encode(result)
	}
}
