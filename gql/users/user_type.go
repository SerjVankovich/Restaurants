package users

import "github.com/graphql-go/graphql"

var HashType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Hash",
		Fields: graphql.Fields{
			"hash": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
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
			"pref_rest": &graphql.Field{
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

var ConfirmedType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Confirmed",
		Fields: graphql.Fields{
			"isOk": &graphql.Field{
				Type: graphql.Boolean,
			},
			"access_token": &graphql.Field{
				Type: graphql.String,
			},
			"confirm_hash": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
