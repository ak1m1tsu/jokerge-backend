package types

import (
	"time"

	"github.com/uptrace/bun"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        string `bun:",pk"`
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type ProductModel struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID          string `bun:",pk"`
	Name        string
	Description string
	Price       int
}

type CustomerModel struct {
	bun.BaseModel `bun:"table:customers,alias:c"`

	ID        string `bun:",pk"`
	FirstName string
	LastName  string
	Address   string
	Orders    []OrderModel `bun:"rel:has-many,join:id=customer_id"`
}

func (m CustomerModel) ToCustomer() *Customer {
	c := &Customer{
		CustomerBase: CustomerBase{
			ID:        m.ID,
			FirstName: m.FirstName,
			LastName:  m.LastName,
			Address:   m.Address,
		},
		Orders: make([]Order, 0),
	}

	return c
}

type OrderModel struct {
	bun.BaseModel `bun:"table:orders,alias:o"`

	ID         int `bun:",pk,autoincrement"`
	CustomerID string
	Customer   *CustomerModel   `bun:"rel:belongs-to,join:customer_id=id"`
	Products   []OrderItemModel `bun:"rel:has-many,join:id=order_id"`
	Status     OrderStatus
	Price      int
	CreatedAt  int64
}

func (m OrderModel) ToOrder() *Order {
	o := &Order{
		OrderBase: OrderBase{
			ID:        m.ID,
			Status:    m.Status,
			Price:     m.Price,
			CreatedAt: time.Unix(m.CreatedAt, 0),
		},
		Products: make([]Product, 0),
	}

	return o
}

func (m OrderModel) ToOrderWithCustomer() *OrderWithCustomer {
	return &OrderWithCustomer{
		Order: Order{
			OrderBase: OrderBase{
				ID:        m.ID,
				Status:    m.Status,
				Price:     m.Price,
				CreatedAt: time.Unix(m.CreatedAt, 0),
			},
			Products: make([]Product, 0),
		},
		Customer: &CustomerBase{
			ID:        m.Customer.ID,
			FirstName: m.Customer.FirstName,
			LastName:  m.Customer.LastName,
			Address:   m.Customer.Address,
		},
	}
}

type OrderItemModel struct {
	bun.BaseModel `bun:"table:order_items,alias:oi"`

	OrderID   int
	Order     *OrderModel `bun:"rel:belongs-to,join:order_id=id"`
	ProductID string
	Product   *ProductModel `bun:"rel:belongs-to,join:product_id=id"`
	Count     int
}
