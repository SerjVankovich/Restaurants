package products

import (
	"../../db"
	"../../models"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
)

func ProductMutationType(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"addProduct": addProduct(dataBase),
			},
		})
}

func addProduct(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: ProductType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"category": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			name, nameOk := p.Args["name"].(string)
			description, descriptionOk := p.Args["description"].(string)
			price, priceOk := p.Args["price"].(float64)
			category, categoryOk := p.Args["category"].(int)

			if !nameOk {
				return nil, errors.New("name not provided")
			}

			if !descriptionOk {
				return nil, errors.New("description not provided")
			}

			if !priceOk {
				return nil, errors.New("price not provided")
			}

			if !categoryOk {
				return nil, errors.New("category not provided")
			}

			product := models.Product{Name: name, Description: description, Price: float32(price), Category: category}

			prod, err := db.AddProduct(dataBase, &product)

			return prod, err
		},
		Description: "Add one product",
	}
}
