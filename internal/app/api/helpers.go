package api

import (
	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
)

func FillCustomerItem(customer types.Customer) CustomerItem {
	return CustomerItem{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Address:   customer.Address,
	}
}

func FillCustomerListItem(customer types.Customer, orders []types.Order) CustomerListItem {
	result := CustomerListItem{
		CustomerItem: FillCustomerItem(customer),
		Orders:       make([]OrderItem, 0, len(orders)),
	}

	for _, order := range orders {
		result.Orders = append(result.Orders, FillOrderItem(order))
	}

	return result
}

func FillCustomerList(customers []types.Customer, orders []types.Order) CustomerList {
	result := make(CustomerList, 0, len(customers))
	customersWithOrders := map[string]*types.CustomerWithOrders{}

	for _, customer := range customers {
		customersWithOrders[customer.ID] = &types.CustomerWithOrders{
			Customer: customer,
			Orders:   make([]types.Order, 0),
		}
	}

	for _, order := range orders {
		if cwo, found := customersWithOrders[order.CustomerID]; found {
			cwo.Orders = append(cwo.Orders, order)
		}
	}

	for _, cwo := range customersWithOrders {
		result = append(result, FillCustomerListItem(cwo.Customer, cwo.Orders))
	}

	return result
}

func FillOrderItem(order types.Order) OrderItem {
	return OrderItem{
		ID:     order.ID,
		Status: order.Status.String(),
	}
}

func FillOrderListItem(order types.Order, customer types.Customer) OrderListItem {
	return OrderListItem{
		OrderItem: FillOrderItem(order),
		Customer:  FillCustomerItem(customer),
	}
}

func FillOrderList(orders []types.Order, customers []types.Customer) OrderList {
	customersWithOrders := map[string]*types.CustomerWithOrders{}
	result := make(OrderList, 0)

	for _, customer := range customers {
		customersWithOrders[customer.ID] = &types.CustomerWithOrders{
			Customer: customer,
			Orders:   make([]types.Order, 0),
		}
	}

	for _, order := range orders {
		if cwo, found := customersWithOrders[order.CustomerID]; found {
			cwo.Orders = append(cwo.Orders, order)
		}
	}

	for _, cwo := range customersWithOrders {
		for _, order := range cwo.Orders {
			result = append(result, FillOrderListItem(order, cwo.Customer))
		}
	}

	return result
}

func FillCustomerOrdersItem(customer types.Customer, orders []types.Order) CustomerOrdersItem {
	result := CustomerOrdersItem{
		Customer: FillCustomerItem(customer),
		Orders:   make(OrderList, 0, len(orders)),
	}
	for _, order := range orders {
		result.Orders = append(result.Orders, FillOrderListItem(order, customer))
	}
	return result
}
