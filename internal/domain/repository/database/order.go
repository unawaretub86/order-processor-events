package database

import "fmt"

const table = "orders"

func (d *databaseOrder) CreateOrder(body, requestId string) error {
	fmt.Println("hi im in database")

	return nil
}
