package orders

import (
	"../../db"
	"database/sql"
	"fmt"
	"github.com/graphql-go/graphql"
	"net/http"
)

func OrderQueryType(dataBase *sql.DB, request *http.Request) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"allOrders": &graphql.Field{
					Type: graphql.NewList(OrderType),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						fmt.Println(p.Context)
						return db.GetAllOrders(dataBase)
					},
					Description: "Get all orders",
				},
				"uncompletedOrders": &graphql.Field{
					Type: graphql.NewList(OrderType),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return db.GetUncompletedOrders(dataBase)
					},
					Description: "Get just uncompleted orders",
				},
				"completedOrders": &graphql.Field{
					Type: graphql.NewList(OrderType),
					Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
						return db.GetCompleteOrders(dataBase)
					},
					Description: "Get just completed orders",
				},
			},
		})
}
