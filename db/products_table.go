package db

import (
	"../models"
	"database/sql"
	"strconv"
)

func getProductsQuery(dataBase *sql.DB, query string) ([]*models.Product, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	rows, err := dataBase.Query(query)

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

func GetAllProducts(dataBase *sql.DB) ([]*models.Product, error) {
	return getProductsQuery(dataBase, "SELECT * FROM products")
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

func GetProductsByCategory(dataBase *sql.DB, category int) ([]*models.Product, error) {
	return getProductsQuery(dataBase, "SELECT * FROM products WHERE category = "+strconv.Itoa(category))
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
