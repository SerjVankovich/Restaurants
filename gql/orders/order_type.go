package orders

import "github.com/graphql-go/graphql"

var OrderItem = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "OrderItem",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"product": &graphql.Field{
				Type: graphql.Int,
			},
			"num_product": &graphql.Field{
				Type: graphql.Int,
			},
			"order": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

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
			"order_items": &graphql.Field{
				Type: graphql.NewList(OrderItem),
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
