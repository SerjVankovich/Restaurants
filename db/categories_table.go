package db

import (
	"../models"
	"database/sql"
)

func GetAllCategories(dataBase *sql.DB) ([]*models.Category, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	var categories []*models.Category

	rows, err := dataBase.Query("SELECT * FROM categories")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category models.Category

		err = rows.Scan(
			&category.Id,
			&category.Name,
			&category.Restaurant,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	return categories, nil
}

func GetCategoriesByRestaurant(dataBase *sql.DB, restaurant int) ([]*models.Category, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	rows, err := dataBase.Query("SELECT * FROM categories WHERE restaurant = $1", restaurant)

	if err != nil {
		return nil, err
	}

	var categories []*models.Category

	for rows.Next() {
		var category models.Category
		err = rows.Scan(
			&category.Id,
			&category.Name,
			&category.Restaurant)

		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	return categories, nil
}

func AddCategory(dataBase *sql.DB, category models.Category) error {
	if dataBase == nil {
		return dbErr
	}

	_, err := dataBase.Exec("INSERT INTO categories (name, restaurant) VALUES ($1, $2)",
		category.Name,
		category.Restaurant)

	return err
}
