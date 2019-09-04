package db

import (
	"../models"
	"database/sql"
)

func getOrdersQuery(dataBase *sql.DB, query string) ([]*models.Order, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	rows, err := dataBase.Query(query)

	if err != nil {
		return nil, err
	}

	var orders []*models.Order

	for rows.Next() {
		var order models.Order

		err = rows.Scan(&order.Id, &order.User, &order.Restaurant, &order.Time, &order.Complete)

		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)

	}

	return orders, nil
}

func GetAllOrders(dataBase *sql.DB) ([]*models.Order, error) {
	return getOrdersQuery(dataBase, "SELECT * FROM orders")
}

func GetUncompletedOrders(dataBase *sql.DB) ([]*models.Order, error) {
	return getOrdersQuery(dataBase, "SELECT * FROM orders WHERE complete = FALSE")
}

func GetCompleteOrders(dataBase *sql.DB) ([]*models.Order, error) {
	return getOrdersQuery(dataBase, "SELECT * FROM orders WHERE complete = TRUE")
}

func AddOrder(dataBase *sql.DB, order *models.Order) (*models.Order, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	_, err := dataBase.Exec(`INSERT INTO orders (id, "user", restaurant, time, complete) VALUES (DEFAULT, $1, $2, $3, DEFAULT)`,
		order.User,
		order.Restaurant,
		order.Time,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func CompleteOrder(dataBase *sql.DB, id int) error {
	if dataBase == nil {
		return dbErr
	}

	_, err := dataBase.Exec(`UPDATE orders SET "complete" = $2 WHERE id = $1`, id, true)

	return err
}
