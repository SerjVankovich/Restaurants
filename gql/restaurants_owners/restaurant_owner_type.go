package restaurants_owners

import "github.com/graphql-go/graphql"

var RestaurantOwnerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RestaurantOwner",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
			"salt": &graphql.Field{
				Type: graphql.String,
			},
			"confirmed": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
