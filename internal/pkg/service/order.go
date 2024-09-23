package service

import (
	"context"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
)

// GetOrders возвращает список заказов с их клиентами
func (s *Service) GetOrders(ctx context.Context) ([]*types.OrderWithCustomer, error) {
	var model []types.OrderModel
	if err := s.db.NewSelect().Model(&model).Scan(ctx); err != nil {
		return nil, err
	}

	orders := make([]*types.OrderWithCustomer, 0, len(model))
	for _, m := range model {
		orders = append(orders, m.ToOrderWithCustomer())
	}

	return orders, nil
}

// GetOrderInfo возвращает информацию о заказе вместе с его клиентом
func (s *Service) GetOrderInfo(ctx context.Context, id string) (*types.OrderWithCustomer, error) {
	order := &types.OrderWithCustomer{}

	return order, nil
}
