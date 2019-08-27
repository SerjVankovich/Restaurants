package restaurants

import (
	"../../db"
	"../../models"
	"database/sql"
	"errors"
	"github.com/graphql-go/graphql"
)

func RestaurantMutation(dataBase *sql.DB) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"addRestaurant": addRestaurant(dataBase),
			},
		})
}

func addRestaurant(dataBase *sql.DB) *graphql.Field {
	return &graphql.Field{
		Type: RestaurantType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"latitude": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"longitude": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"owner": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			name, nameOk := p.Args["name"].(string)
			description, descriptionOk := p.Args["description"].(string)
			latitude, latitudeOk := p.Args["latitude"].(float64)
			longitude, longitudeOk := p.Args["longitude"].(float64)
			owner, ownerOk := p.Args["owner"].(int)

			if !nameOk {
				return nil, errors.New("name not provided")
			}

			if !descriptionOk {
				return nil, errors.New("description not provided")
			}

			if !latitudeOk {
				return nil, errors.New("latitude not provided")
			}

			if !longitudeOk {
				return nil, errors.New("longitude not provided")
			}

			if !ownerOk {
				return nil, errors.New("owner not provided")
			}

			restaurant := models.Restaurant{Name: name, Description: description, Latitude: float32(latitude), Longitude: float32(longitude), Owner: int32(owner)}

			err := db.AddRestaurant(dataBase, &restaurant)

			if err != nil {
				return nil, err
			}

			return restaurant, nil

		},
		Description: "Add restaurant",
	}
}
