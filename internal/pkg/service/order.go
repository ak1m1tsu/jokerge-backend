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
			OrderID:   "87aa13fc-81f8-4712-bb38-4415eb2e294d",
			Status:    types.OrderStatusActive,
			CreatedAt: time.Now(),
		},
		{
			OrderID:   "f9593790-4948-4cb4-8ca0-1510bb3ef116",
			Status:    types.OrderStatusCompleted,
			CreatedAt: time.Now(),
		},
		{
			OrderID:   "bd512907-11ec-402d-a05a-7c9d147e23ef",
			Status:    types.OrderStatusActive,
			CreatedAt: time.Now(),
		},
		{
			OrderID:   "b35248c2-559c-4bb1-8d43-9b1aaa0df8f4",
			Status:    types.OrderStatusActive,
			CreatedAt: time.Now(),
		},
		{
			OrderID:   "2bb7d14a-d98a-41ce-a0bc-764798672776",
			Status:    types.OrderStatusCanceled,
			CreatedAt: time.Now(),
		},
		{
			OrderID:   "cfd373d6-4708-461d-892a-327b6dae27d7",
			Status:    types.OrderStatusActive,
			CreatedAt: time.Now(),
		},
	}

	return orders, nil
}

func (s *Service) GetOrderInfo(ctx context.Context, id string) (*types.Order, error) {
	order := &types.Order{
		OrderID:   "2bb7d14a-d98a-41ce-a0bc-764798672776",
		Status:    types.OrderStatusCanceled,
		CreatedAt: time.Now(),
	}

	return order, nil
}
