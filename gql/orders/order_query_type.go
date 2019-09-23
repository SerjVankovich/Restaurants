package orders

import (
	"../../db"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
	"net/http"
	"restaurants/utils"
)

func OrderQueryType(dataBase *sql.DB, request *http.Request) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"allOrders": &graphql.Field{
					Type:        graphql.NewList(OrderType),
					Resolve:     allOrders(dataBase, request),
					Description: "Get all orders",
				},
				"uncompletedOrders": &graphql.Field{
					Type:        graphql.NewList(OrderType),
					Resolve:     uncompletedOrders(dataBase, request),
					Description: "Get just uncompleted orders",
				},
				"completedOrders": &graphql.Field{
					Type:        graphql.NewList(OrderType),
					Resolve:     completedOrders(dataBase, request),
					Description: "Get just completed orders",
				},
				"ordersByRestaurant": &graphql.Field{
					Type: graphql.NewList(OrderType),
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve:     ordersByRestaurant(dataBase, request),
					Description: "Get all orders by restaurant",
				},
				"ordersByUser": &graphql.Field{
					Type: graphql.NewList(OrderType),
					Args: graphql.FieldConfigArgument{
						"user": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve:     ordersByUser(dataBase, request),
					Description: "Get orders by user",
				},
			},
		})
}

func ordersByUser(dataBase *sql.DB, request *http.Request) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		user, userOk := p.Args["user"].(int)

		if !userOk {
			return nil, errors.New("user id not provided")
		}

		err := utils.SimpleValidateToken(request.Header.Get(utils.AUTHORIZATION), utils.USER)

		if err != nil {
			return nil, err
		}

		return db.GetOrdersByUser(dataBase, user)
	}
}

func ordersByRestaurant(dataBase *sql.DB, request *http.Request) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {
		id, idOk := p.Args["id"].(int)

		if !idOk {
			return nil, errors.New("restaurant id not provided")
		}

		err := utils.SimpleValidateToken(request.Header.Get(utils.AUTHORIZATION), utils.OWNER)

		if err != nil {
			return nil, err
		}

		return db.GetOrdersByRestaurant(dataBase, id)
	}
}

func completedOrders(dataBase *sql.DB, request *http.Request) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {

		err := utils.SimpleValidateToken(request.Header.Get(utils.AUTHORIZATION), utils.OWNER)

		if err != nil {
			return nil, err
		}

		return db.GetCompleteOrders(dataBase)
	}
}

func uncompletedOrders(dataBase *sql.DB, request *http.Request) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, e error) {

		err := utils.SimpleValidateToken(request.Header.Get(utils.AUTHORIZATION), utils.OWNER)

		if err != nil {
			return nil, err
		}

		return db.GetUncompletedOrders(dataBase)
	}
}

func allOrders(dataBase *sql.DB, request *http.Request) func(p graphql.ResolveParams) (i interface{}, e error) {
	return func(p graphql.ResolveParams) (i interface{}, e error) {

		err := utils.SimpleValidateToken(request.Header.Get(utils.AUTHORIZATION), utils.OWNER)

		if err != nil {
			return nil, err
		}

		return db.GetAllOrders(dataBase)
	}
}
