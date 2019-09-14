package orders

import (
	"../../db"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
	"net/http"
	"restaurants/utils"
)

func queryValidate(header string) error {
	claims, err := utils.ValidateToken(header)
	if err != nil {
		return err
	}
	userType := claims["type"].(string)
	if userType != "owner" {
		return errors.New(`user type is not "owner" but ` + userType)
	}

	return nil
}

func OrderQueryType(dataBase *sql.DB, request *http.Request) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"allOrders": &graphql.Field{
					Type: graphql.NewList(OrderType),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

						err := queryValidate(request.Header.Get("Authorization"))

						if err != nil {
							return nil, err
						}

						return db.GetAllOrders(dataBase)
					},
					Description: "Get all orders",
				},
				"uncompletedOrders": &graphql.Field{
					Type: graphql.NewList(OrderType),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

						err := queryValidate(request.Header.Get("Authorization"))

						if err != nil {
							return nil, err
						}
						return db.GetUncompletedOrders(dataBase)
					},
					Description: "Get just uncompleted orders",
				},
				"completedOrders": &graphql.Field{
					Type: graphql.NewList(OrderType),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

						err := queryValidate(request.Header.Get("Authorization"))

						if err != nil {
							return nil, err
						}
						return db.GetCompleteOrders(dataBase)
					},
					Description: "Get just completed orders",
				},
				"ordersByRestaurant": &graphql.Field{
					Type: graphql.NewList(OrderType),
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						id, idOk := p.Args["id"].(int)

						if !idOk {
							return nil, errors.New("restaurant id not provided")
						}

						err := queryValidate(request.Header.Get("Authorization"))

						if err != nil {
							return nil, err
						}

						return db.GetOrdersByRestaurant(dataBase, id)
					},
					Description: "Get all orders by restaurant",
				},
			},
		})
}
