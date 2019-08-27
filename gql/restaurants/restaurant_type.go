package restaurants

import "github.com/graphql-go/graphql"

var RestaurantType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RestaurantType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"latitude": &graphql.Field{
				Type: graphql.Float,
			},
			"longitude": &graphql.Field{
				Type: graphql.Float,
			},
			"owner": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
