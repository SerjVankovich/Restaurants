package products

import (
	"../../db"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
)

func ProductQueryType(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"products":           products(dataBase),
				"productsByCategory": productsByCategory(dataBase),
				"productsById":       productsById(dataBase),
			},
		})
}

func productsById(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: ProductType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id, idOk := p.Args["id"].(int)

			if !idOk {
				return nil, errors.New("id not provided")
			}

			return db.GetProductById(dataBase, id)
		},
		Description: "Get product by Id",
	}
}

func productsByCategory(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(ProductType),
		Args: graphql.FieldConfigArgument{
			"category": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			category, categoryOk := p.Args["category"].(int)

			if !categoryOk {
				return nil, errors.New("category not provided")
			}

			return db.GetProductsByCategory(dataBase, category)
		},
		Description: "Get all products by category",
	}
}

func products(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(ProductType),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return db.GetAllProducts(dataBase)
		},
		Description: "Get all products",
	}
}
