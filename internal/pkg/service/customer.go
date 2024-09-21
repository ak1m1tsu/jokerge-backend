package service

import (
	"context"
	"time"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
)

// GetCustomerInfo возвращает информацию о клиенте
func (s *Service) GetCustomerInfo(ctx context.Context, customerID string) (*types.Customer, error) {
	return &types.Customer{
		ID:        "7dfa4d19-d4c8-4596-bcb7-249e0de951bc",
		FirstName: "Инесса",
		LastName:  "Химченко",
		Address:   "164505, г. Дятьково, ул. Белоостровская (Выборгский), дом 59, квартира 24",
	}, nil
}

// GetCustomerOrders возвращает заказы клиента
func (s *Service) GetCustomerOrders(ctx context.Context, customerID string) ([]types.Order, error) {
	return []types.Order{
		{
			ID:         "2bb7d14a-d98a-41ce-a0bc-764798672776",
			CustomerID: "7dfa4d19-d4c8-4596-bcb7-249e0de951bc",
			Status:     types.OrderStatusCanceled,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "bd512907-11ec-402d-a05a-7c9d147e23ef",
			CustomerID: "7dfa4d19-d4c8-4596-bcb7-249e0de951bc",
			Status:     types.OrderStatusActive,
			CreatedAt:  time.Now(),
		},
	}, nil
}
