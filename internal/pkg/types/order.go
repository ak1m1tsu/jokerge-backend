package types

import "time"

// Order заказ клиента
type Order struct {
	OrderID   string     `bun:",pk"`
	Customer  *Customer  `bun:"rel:belongs-to"`
	Products  OrderItems `bun:"rel:has-many"`
	Status    OrderStatus
	Price     int
	CreatedAt time.Time
}

type OrderItems []OrderItem

type OrderItem struct {
	Product
	Count int
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
