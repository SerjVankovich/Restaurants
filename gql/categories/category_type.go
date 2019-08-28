package categories

import "github.com/graphql-go/graphql"

var CategoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"restaurant": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
