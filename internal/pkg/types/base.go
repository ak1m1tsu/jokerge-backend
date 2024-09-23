package types

import "time"

// CustomerBase базовая презентация клиента
type CustomerBase struct {
	ID        string
	FirstName string
	LastName  string
	Address   string
}

// OrderBase базовая презентация заказа
type OrderBase struct {
	ID        int
	Status    OrderStatus
	Price     int
	CreatedAt time.Time
}

// ProdcutBase базовая презентация продукта
type ProductBase struct {
	ID          string
	Name        string
	Description string
	Price       int
}

// UserBase базовая презентация пользователя
type UserBase struct {
	ID        string
	Email     string
	Password  string
	FirstName string
	LastName  string
}
