package types

// Customer пользователь, который оформляет заказы
type Customer struct {
	ID        string `bun:",pk"`
	FirstName string
	LastName  string
	Address   string
	Orders    []Order `bun:"rel:has-many"`
}

// Client пользователь и его история заказов
type CustomerWithOrders struct {
	Customer `bun:"rel:belongs-to"`
	Orders   []Order `bun:"rel:has-many"`
}
