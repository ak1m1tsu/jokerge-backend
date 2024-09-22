package types

import "time"

// Order заказ клиента
type Order struct {
	ID         string `bun:",pk"`
	CustomerID string
	Status     OrderStatus
	CreatedAt  time.Time
}

// OrderStatus статус заказа
type OrderStatus int

func (os OrderStatus) String() string {
	return orderStatuses[os]
}

const (
	OrderStatusActive OrderStatus = iota
	OrderStatusCompleted
	OrderStatusCanceled
)

var orderStatuses = map[OrderStatus]string{
	OrderStatusActive:    "active",
	OrderStatusCompleted: "completed",
	OrderStatusCanceled:  "canceled",
}
