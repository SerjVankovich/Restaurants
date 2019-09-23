package db

import (
	"../models"
	"errors"
	"strconv"
)

func GetAllItemsByOrder(dataBase dbInterface, order int) ([]*models.OrderItem, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	rows, err := dataBase.Query(`SELECT * FROM order_items WHERE "order" = $1`, order)

	if err != nil {
		return nil, err
	}

	var orderItems []*models.OrderItem

	for rows.Next() {
		var item models.OrderItem

		err = rows.Scan(&item.Id, &item.Product, &item.NumProduct, &item.Order)

		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, &item)
	}

	return orderItems, nil

}

func AddItems(dataBase dbInterface, orderItems []*models.OrderItem, id int) error {
	if dataBase == nil {
		return dbErr
	}

	if orderItems == nil {
		return errors.New("items not provided")
	}

	query := `INSERT INTO order_items (id, product, num_product, "order") values `
	for _, item := range orderItems {
		query += `(DEFAULT, ` + strconv.Itoa(item.Product) + `, ` + strconv.Itoa(item.NumProduct) + `, ` +
			strconv.Itoa(id) + `), `
	}

	_, err := dataBase.Exec(query[:len(query)-2])

	return err
}
