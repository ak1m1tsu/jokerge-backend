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
	Orders      []OrderModel `bun:"m2m:order_items,join:Product=Order"`
}

func (m ProductModel) ToProduct() *Product {
	return &Product{
		ProductBase: ProductBase{
			ID:          m.ID,
			Name:        m.Name,
			Description: m.Description,
			Price:       m.Price,
		},
	}
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

	for _, om := range m.Orders {
		c.Orders = append(c.Orders, *om.ToOrder())
	}

	return c
}

type OrderModel struct {
	bun.BaseModel `bun:"table:orders,alias:o"`

	ID         int `bun:",pk,autoincrement"`
	CustomerID string
	Customer   *CustomerModel   `bun:"rel:belongs-to,join:customer_id=id"`
	Products   []OrderItemModel `bun:"m2m:order_items,join:Order=Product"`
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
		Products: make([]OrderItem, 0),
	}

	for _, pm := range m.Products {
		o.Products = append(o.Products, *pm.ToOrderItem())
	}

	return o
}

func (m OrderModel) ToOrderWithCustomer() *OrderWithCustomer {
	o := &OrderWithCustomer{
		Order: Order{
			OrderBase: OrderBase{
				ID:        m.ID,
				Status:    m.Status,
				Price:     m.Price,
				CreatedAt: time.Unix(m.CreatedAt, 0),
			},
			Products: make([]OrderItem, 0),
		},
	}

	if m.Customer != nil {
		o.Customer = &CustomerBase{
			ID:        m.Customer.ID,
			FirstName: m.Customer.FirstName,
			LastName:  m.Customer.LastName,
			Address:   m.Customer.Address,
		}
	}

	for _, pm := range m.Products {
		o.Products = append(o.Products, *pm.ToOrderItem())
	}

	return o
}

type OrderItemModel struct {
	bun.BaseModel `bun:"table:order_items,alias:oi"`

	OrderID   int           `bun:",pk"`
	Order     *OrderModel   `bun:"rel:belongs-to,join:order_id=id"`
	ProductID string        `bun:",pk"`
	Product   *ProductModel `bun:"rel:belongs-to,join:product_id=id"`
	Count     int
}

func (m OrderItemModel) ToOrderItem() *OrderItem {
	return &OrderItem{
		Product: *m.Product.ToProduct(),
		Count:   m.Count,
	}
}
