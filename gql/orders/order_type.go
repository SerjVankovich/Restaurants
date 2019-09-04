package orders

import "github.com/graphql-go/graphql"

var OrderType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"user": &graphql.Field{
				Type: graphql.Int,
			},
			"restaurant": &graphql.Field{
				Type: graphql.Int,
			},
			"time": &graphql.Field{
				Type: graphql.DateTime,
			},
			"complete": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})

var CompleteOk = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CompleteOk",
		Fields: graphql.Fields{
			"completed": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})
