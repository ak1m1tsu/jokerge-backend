package types

// Customer пользователь, который оформляет заказы
type Customer struct {
	ID        string `json:"id" bun:",pk"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

// Client пользователь и его история заказов
type CustomerWithOrders struct {
	Customer `json:"client" bun:"rel:belongs-to,join:customer_id=id"`
	Orders   []Order `json:"orders" bun:"rel:has-many,join:id=order_id"`
}
