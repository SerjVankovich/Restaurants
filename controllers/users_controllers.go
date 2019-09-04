package controllers

import (
	"database/sql"
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

func GQLHandlerPost(dataBase *sql.DB, schema func(*sql.DB, *http.Request) (graphql.Schema, error)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		sch, err := schema(dataBase, request)

		if err != nil {
			panic(err)
		}

		result := executeQuery(request.URL.Query().Get("query"), sch)
		_ = json.NewEncoder(writer).Encode(result)
	}

}

func GQLHandler(schema graphql.Schema) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		result := executeQuery(request.URL.Query().Get("query"), schema)
		_ = json.NewEncoder(writer).Encode(result)
	}
}
