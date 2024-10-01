package service

import (
	"context"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	"github.com/google/uuid"
)

// GetCustomerInfo возвращает информацию о клиенте и его заказах
func (s *Service) GetCustomerInfo(ctx context.Context, customerID string) (*types.Customer, error) {
	model := new(types.CustomerModel)
	if err := s.db.NewSelect().Model(model).Relation("Orders").Where("c.id = ?", customerID).Scan(ctx); err != nil {
		return nil, err
	}

	return model.ToCustomer(), nil
}

// GetCustomers возвращает список клиентов с их заказами
func (s *Service) GetCustomers(ctx context.Context) ([]*types.Customer, error) {
	var model []types.CustomerModel
	if err := s.db.NewSelect().Model(&model).Relation("Orders").Scan(ctx); err != nil {
		return nil, err
	}

	customers := make([]*types.Customer, 0, len(model))
	for _, m := range model {
		customers = append(customers, m.ToCustomer())
	}

	return customers, nil
}

// CreateCustomer создает нового клиента
func (s *Service) CreateCustomer(ctx context.Context, body types.CustomerCreateBody) (string, error) {
	model := types.CustomerModel{
		ID:        uuid.NewString(),
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Address:   body.Address,
	}
	if _, err := s.db.NewInsert().Model(&model).Exec(ctx); err != nil {
		return "", err
	}

	return model.ID, nil
}

// UpdateCustomer обновляет клиента
func (s *Service) UpdateCustomer(ctx context.Context, body types.CustomerUpdateBody) error {
	model := types.CustomerModel{
		ID:        body.ID,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Address:   body.Address,
	}

	_, err := s.db.NewUpdate().Model(&model).OmitZero().WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
