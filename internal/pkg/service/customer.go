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

// GetCustomers возвращает список клиентов
func (s *Service) GetCustomers(ctx context.Context) ([]types.Customer, error) {
	customers := []types.Customer{
		{
			ID:        "c22946d7-991e-44a1-b0dc-6b775a34664c",
			FirstName: "Модест",
			LastName:  "Пономарёв",
			Address:   "121248, г. Приозерное, ул. Пяловская, дом 40, квартира 30",
		},
		{
			ID:        "4c8a6931-04cb-4c93-9a90-e7a29342aba2",
			FirstName: "Мстислав",
			LastName:  "Рощин",
			Address:   "396522, г. Надеждино, ул. Жилой поселок 2-й Заельцовского Бора тер, дом 41, квартира 96",
		},
		{
			ID:        "7dfa4d19-d4c8-4596-bcb7-249e0de951bc",
			FirstName: "Инесса",
			LastName:  "Химченко",
			Address:   "164505, г. Дятьково, ул. Белоостровская (Выборгский), дом 59, квартира 24",
		},
	}

	return customers, nil
}
