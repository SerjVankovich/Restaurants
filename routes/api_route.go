package routes

import (
	"../controllers"
	"../gql/categories"
	"../gql/products"
	"../gql/restaurants"
	"../gql/users"
	"database/sql"
	"github.com/gorilla/mux"
)

func ApiRoute(dataBase *sql.DB) *mux.Router {
	userSchema, err := users.UserSchema(dataBase)
	if err != nil {
		panic(err)
	}

	restaurantsSchema, err := restaurants.RestaurantSchema(dataBase)

	if err != nil {
		panic(err)
	}

	categoriesSchema, err := categories.CategorySchema(dataBase)

	if err != nil {
		panic(err)
	}

	productsSchema, err := products.ProductSchema(dataBase)

	if err != nil {
		panic(err)
	}
	route := "/api/v1/"
	r := mux.NewRouter()
	r.HandleFunc(route+"restaurants", controllers.GQLHandler(restaurantsSchema))
	r.HandleFunc(route+"users", controllers.GQLHandler(userSchema))
	r.HandleFunc(route+"restaurants_owner", controllers.MockRestaurantsOwnerController)
	r.HandleFunc(route+"categories", controllers.GQLHandler(categoriesSchema))
	r.HandleFunc(route+"products", controllers.GQLHandler(productsSchema))

	return r
}
