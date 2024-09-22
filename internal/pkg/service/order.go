package service

import (
	"context"
	"time"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
)

// GetOrders возвращает список заказов
func (s *Service) GetOrders(ctx context.Context) ([]types.Order, error) {
	orders := []types.Order{
		{
			ID:         "87aa13fc-81f8-4712-bb38-4415eb2e294d",
			CustomerID: "c22946d7-991e-44a1-b0dc-6b775a34664c",
			Status:     types.OrderStatusActive,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "f9593790-4948-4cb4-8ca0-1510bb3ef116",
			CustomerID: "c22946d7-991e-44a1-b0dc-6b775a34664c",
			Status:     types.OrderStatusCompleted,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "bd512907-11ec-402d-a05a-7c9d147e23ef",
			CustomerID: "7dfa4d19-d4c8-4596-bcb7-249e0de951bc",
			Status:     types.OrderStatusActive,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "b35248c2-559c-4bb1-8d43-9b1aaa0df8f4",
			CustomerID: "4c8a6931-04cb-4c93-9a90-e7a29342aba2",
			Status:     types.OrderStatusActive,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "2bb7d14a-d98a-41ce-a0bc-764798672776",
			CustomerID: "7dfa4d19-d4c8-4596-bcb7-249e0de951bc",
			Status:     types.OrderStatusCanceled,
			CreatedAt:  time.Now(),
		},
		{
			ID:         "cfd373d6-4708-461d-892a-327b6dae27d7",
			CustomerID: "4c8a6931-04cb-4c93-9a90-e7a29342aba2",
			Status:     types.OrderStatusActive,
			CreatedAt:  time.Now(),
		},
	}

	return orders, nil
}

func (s *Service) GetOrderInfo(ctx context.Context, id string) (*types.Order, error) {
	order := &types.Order{
		ID:         "2bb7d14a-d98a-41ce-a0bc-764798672776",
		CustomerID: "7dfa4d19-d4c8-4596-bcb7-249e0de951bc",
		Status:     types.OrderStatusCanceled,
		CreatedAt:  time.Now(),
	}

	return order, nil
}
