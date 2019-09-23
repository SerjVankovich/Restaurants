package tests

import (
	"database/sql"
	"restaurants/models"
	"testing"
)

type dbMockOrders struct {
	orders []*models.Order
}

func (*dbMockOrders) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// TODO make rows interface to test
	return nil, nil
}
func TestGetOrdersQuery(t *testing.T) {

}
