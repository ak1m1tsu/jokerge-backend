package service

import (
	"context"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
)

// GetCustomerInfo возвращает информацию о клиенте и его заказах
func (s *Service) GetCustomerInfo(ctx context.Context, customerID string) (*types.Customer, error) {
	model := new(types.CustomerModel)
	if err := s.db.NewSelect().Model(model).Where("id = ?", customerID).Scan(ctx); err != nil {
		return nil, err
	}

	return model.ToCustomer(), nil
}

// GetCustomers возвращает список клиентов с их заказами
func (s *Service) GetCustomers(ctx context.Context) ([]*types.Customer, error) {
	var model []types.CustomerModel
	if err := s.db.NewSelect().Model(&model).Scan(ctx); err != nil {
		return nil, err
	}

	customers := make([]*types.Customer, 0, len(model))
	for _, m := range model {
		customers = append(customers, m.ToCustomer())
	}

	return customers, nil
}
