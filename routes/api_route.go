package routes

import (
	"../controllers"
	"../gql/categories"
	"../gql/orders"
	"../gql/products"
	"../gql/restaurants"
	"../gql/restaurants_owners"
	"../gql/users"
	"database/sql"
	"github.com/gorilla/mux"
)

func ApiRoute(dataBase *sql.DB) *mux.Router {
	userSchema, err := users.UserSchema(dataBase)

	if err != nil {
		panic(err)
	}

	restaurantOwnerSchema, err := restaurants_owners.RestaurantOwnerSchema(dataBase)

	if err != nil {
		panic(err)
	}

	route := "/api/v1/"
	r := mux.NewRouter()
	r.HandleFunc(route+"restaurants", controllers.GQLHandlerWithRequest(dataBase, restaurants.RestaurantSchema))
	r.HandleFunc(route+"users", controllers.GQLHandler(userSchema))
	r.HandleFunc(route+"restaurants_owners", controllers.GQLHandler(restaurantOwnerSchema))
	r.HandleFunc(route+"categories", controllers.GQLHandlerWithRequest(dataBase, categories.CategorySchema))
	r.HandleFunc(route+"products", controllers.GQLHandlerWithRequest(dataBase, products.ProductSchema))
	r.HandleFunc(route+"orders", controllers.GQLHandlerWithRequest(dataBase, orders.OrderSchema))

	return r
}
