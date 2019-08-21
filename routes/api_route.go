package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)

func ApiRoute() *mux.Router {
	route := "/api/v1/"
	r := mux.NewRouter()
	r.HandleFunc(route+"restaurants", controllers.MockRestaurantsController)
	r.HandleFunc(route+"users", controllers.MockUsersController)
	r.HandleFunc(route+"restaurants_owner", controllers.MockRestaurantsOwnerController)

	return r
}
