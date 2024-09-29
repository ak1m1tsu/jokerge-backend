package service

import (
	"context"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	"github.com/rs/zerolog"
)

// GetOrders возвращает список заказов с их клиентами
func (s *Service) GetOrders(ctx context.Context) ([]*types.OrderWithCustomer, error) {
	var model []types.OrderModel
	if err := s.db.NewSelect().Model(&model).Relation("Customer").Scan(ctx); err != nil {
		return nil, err
	}

	orders := make([]*types.OrderWithCustomer, 0, len(model))
	for _, m := range model {
		if err := s.db.NewSelect().
			Model(&m.Products).
			Relation("Product").
			Where("oi.order_id = ?", m.ID).
			Scan(ctx); err != nil {
			return nil, err
		}

		zerolog.Ctx(ctx).Info().Any("products", m.Products).Int("count", len(m.Products)).Msg("products")

		orders = append(orders, m.ToOrderWithCustomer())
	}

	return orders, nil
}

// GetOrderInfo возвращает информацию о заказе вместе с его клиентом
func (s *Service) GetOrderInfo(ctx context.Context, id int) (*types.OrderWithCustomer, error) {
	model := new(types.OrderModel)
	if err := s.db.NewSelect().Model(model).Where("o.id = ?", id).Relation("Customer").Scan(ctx); err != nil {
		return nil, err
	}

	if err := s.db.NewSelect().Model(&model.Products).Where("oi.order_id = ?", model.ID).Relation("Product").Scan(ctx); err != nil {
		return nil, err
	}

	return model.ToOrderWithCustomer(), nil
}
