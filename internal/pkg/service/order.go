package service

import (
	"context"
	"time"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	zlog "github.com/rs/zerolog"
	"github.com/uptrace/bun"
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

		zlog.Ctx(ctx).Info().Any("products", m.Products).Int("count", len(m.Products)).Msg("products")

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

// CreateOrder создает новый заказ
func (s *Service) CreateOrder(ctx context.Context, body types.CreateOrderBody) (int, error) {
	model := types.OrderModel{
		CustomerID: body.CustomerID,
		Status:     types.OrderStatusActive,
		CreatedAt:  time.Now().Unix(),
		Products:   make([]types.OrderItemModel, 0),
	}

	productIDs := make([]string, 0)
	for id := range body.Products {
		productIDs = append(productIDs, id)
	}

	pmodel := []types.ProductModel{}
	if err := s.db.NewSelect().Model(&pmodel).Where("id IN (?)", bun.In(productIDs)).Scan(ctx); err != nil {
		return 0, err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	zlog.Ctx(ctx).Info().Any("products", pmodel).Any("ids", productIDs).Msg("found products")

	for _, p := range pmodel {
		model.Price = p.Price * body.Products[p.ID]
	}

	if _, err = tx.NewInsert().Model(&model).Exec(ctx); err != nil {
		return 0, err
	}

	zlog.Ctx(ctx).Info().Any("order", model).Msg("order had been created")

	var items []types.OrderItemModel
	for _, p := range pmodel {
		items = append(items, types.OrderItemModel{
			OrderID:   model.ID,
			ProductID: p.ID,
			Count:     body.Products[p.ID],
		})
	}

	zlog.Ctx(ctx).Info().Any("order_items", items).Msg("order items")

	if _, err = tx.NewInsert().Model(&items).Exec(ctx); err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, nil
	}

	return model.ID, nil
}
