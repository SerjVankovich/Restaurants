package db

import (
	"../models"
	"database/sql"
)

func GetAllRestaurants(dataBase *sql.DB) ([]*models.Restaurant, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	rows, err := dataBase.Query("SELECT * FROM restaurants")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var restaurants []*models.Restaurant

	for rows.Next() {
		var restaurant models.Restaurant

		err := rows.Scan(
			&restaurant.Id,
			&restaurant.Name,
			&restaurant.Latitude,
			&restaurant.Longitude,
			&restaurant.Description,
			&restaurant.Owner)

		if err != nil {
			return nil, err
		}

		restaurants = append(restaurants, &restaurant)
	}

	return restaurants, nil
}

func AddRestaurant(dataBase *sql.DB, restaurant *models.Restaurant) error {
	_, err := dataBase.Exec("INSERT INTO restaurants (id, name, latitude, longitude, description, owner)"+
		" values (DEFAULT, $1, $2, $3, $4, $5)",
		restaurant.Name,
		restaurant.Latitude,
		restaurant.Longitude,
		restaurant.Description,
		restaurant.Owner)

	return err
}
