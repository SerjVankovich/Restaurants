package db

import (
	"../models"
	"database/sql"
)

func GetAllProducts(dataBase *sql.DB) ([]*models.Product, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	rows, err := dataBase.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	var products []*models.Product

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Category)

		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func GetProductById(dataBase *sql.DB, id int) (*models.Product, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	row := dataBase.QueryRow("SELECT * FROM products WHERE id = $1", id)

	var product models.Product

	err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Category)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func GetProductByCategory(dataBase *sql.DB, category int) ([]*models.Product, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	rows, err := dataBase.Query("SELECT * FROM products WHERE category = $1", category)

	if err != nil {
		return nil, err
	}

	var products []*models.Product

	for rows.Next() {
		var product models.Product

		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Category)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil

}

func AddProduct(dataBase *sql.DB, product *models.Product) (*models.Product, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	_, err := dataBase.Exec("INSERT INTO products (id, name, description, price, category) VALUES (DEFAULT, $1, $2, $3, $4)",
		product.Name,
		product.Description,
		product.Price,
		product.Category,
	)

	if err != nil {
		return nil, err
	}

	return product, nil

}
