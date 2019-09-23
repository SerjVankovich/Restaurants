package db

import (
	"database/sql"
	"restaurants/models"
	"strconv"
)

func getOrdersQuery(dataBase dbInterface, query string) ([]*models.Order, error) {
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

		items, err := GetAllItemsByOrder(dataBase, order.Id)

		if err != nil {
			return nil, err
		}

		order.OrderItems = items

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

func GetOrdersByRestaurant(dataBase *sql.DB, id int) ([]*models.Order, error) {
	return getOrdersQuery(dataBase, "SELECT * FROM orders WHERE restaurant = "+strconv.Itoa(id))
}

func GetOrdersByUser(dataBase *sql.DB, user int) ([]*models.Order, error) {
	return getOrdersQuery(dataBase, "SELECT * FROM orders WHERE user = "+strconv.Itoa(user))
}

func AddOrder(dataBase *sql.DB, order *models.Order) (*models.Order, error) {
	if dataBase == nil {
		return nil, dbErr
	}

	row := dataBase.QueryRow(`INSERT INTO orders (id, "user", restaurant, time, complete) VALUES (DEFAULT, $1, $2, $3, DEFAULT) RETURNING id`,
		order.User,
		order.Restaurant,
		order.Time,
	)

	var id int

	err := row.Scan(&id)

	if err != nil {
		return nil, err
	}
	order.Id = id

	return order, nil
}

func CompleteOrder(dataBase *sql.DB, id int) error {
	if dataBase == nil {
		return dbErr
	}

	_, err := dataBase.Exec(`UPDATE orders SET "complete" = $2 WHERE id = $1`, id, true)

	return err
}
