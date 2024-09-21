package types

// Customer пользователь, который оформляет заказы
type Customer struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
}

// Client пользователь и его история заказов
type CustomerWithOrders struct {
	Customer `json:"client"`
	Orders   []Order `json:"orders"`
}
