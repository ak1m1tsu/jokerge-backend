package types

// Customer полноценная презенация заказа
type Customer struct {
	CustomerBase
	Orders []Order
}

// Order полноценная презентация заказа
type Order struct {
	OrderBase
	Products []OrderItem
}

// CalculatePrice считает текущую стоимость заказа, исходя из стоимости и кол-ва продуктов.
func (o *Order) CalculatePrice() int {
	var price int
	for _, p := range o.Products {
		price += p.Price * p.Count
	}
	return price
}

// ActualizePrice подсчитывает стоимость заказа исходя из стоимости и кол-ва продуктов,
// и если текущая цена заказа не совпадает с подсчитаной, то старая заменяется новой.
func (o *Order) ActualizePrice() {
	var price = o.CalculatePrice()
	if o.Price != price {
		o.Price = price
	}
}

// OrderWithCustomer представление заказа с клиентом
type OrderWithCustomer struct {
	Order
	Customer *CustomerBase
}

// Product полноценная презентация продукта
type Product struct {
	ProductBase
}

type OrderItem struct {
	Product
	Count int
}

// User полноценная презентация пользователя
type User struct {
	UserBase
}
