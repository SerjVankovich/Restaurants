package routes

import (
	"../controllers"
	"../gql/users"
	"database/sql"
	"github.com/gorilla/mux"
)

func ApiRoute(dataBase *sql.DB) *mux.Router {
	userSchema, err := users.UserSchema(dataBase)
	if err != nil {
		panic(err)
	}
	route := "/api/v1/"
	r := mux.NewRouter()
	r.HandleFunc(route+"restaurants", controllers.MockRestaurantsController)
	r.HandleFunc(route+"users", controllers.GQLHandler(userSchema))
	r.HandleFunc(route+"restaurants_owner", controllers.MockRestaurantsOwnerController)

	return r
}
