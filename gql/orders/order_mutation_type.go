package orders

import (
	"../../db"
	"../../models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"time"
)

func OrderMutationType(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"addOrder":      addOrder(dataBase),
				"completeOrder": completeOrder(dataBase),
			},
		})
}

func completeOrder(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: CompleteOk,
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

			err := db.CompleteOrder(dataBase, id)

			fmt.Println(err)

			return map[string]bool{"completed": true}, err
		},
		Description: "Complete one order",
	}
}

func addOrder(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: OrderType,
		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"restaurant": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			user, userOk := p.Args["user"].(int)
			restaurant, restaurantOk := p.Args["restaurant"].(int)

			if !restaurantOk {
				return nil, errors.New("restaurant not provided")
			}

			if !userOk {
				return nil, errors.New("user not provided")
			}

			order := models.Order{User: user, Restaurant: restaurant, Time: time.Now().Local()}

			ord, err := db.AddOrder(dataBase, &order)

			return ord, err
		},
		Description: "Add one order",
	}
}
