package service

import (
	"context"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	"github.com/google/uuid"
)

// GetProducts возвращает список продуктов
func (s *Service) GetProducts(ctx context.Context) ([]*types.Product, error) {
	var model []types.ProductModel
	if err := s.db.NewSelect().Model(&model).Scan(ctx); err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0, len(model))
	for _, m := range model {
		products = append(products, m.ToProduct())
	}

	return products, nil
}

// GetProductInfo возварщает информацию о продукте
func (s *Service) GetProductInfo(ctx context.Context, id string) (*types.Product, error) {
	model := new(types.ProductModel)
	if err := s.db.NewSelect().Model(model).Where("p.id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}

	return model.ToProduct(), nil
}

// CreateProduct создает новый продукт
func (s *Service) CreateProduct(ctx context.Context, body types.ProductCreateBody) (string, error) {
	model := types.ProductModel{
		ID:          uuid.NewString(),
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
	}

	if _, err := s.db.NewInsert().Model(&model).Exec(ctx); err != nil {
		return "", err
	}

	return model.ID, nil
}

// UpdateProduct обновляет продукт
func (s *Service) UpdateProduct(ctx context.Context, body types.ProductUpdateBody) error {
	model := types.ProductModel{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
	}

	_, err := s.db.NewUpdate().Model(&model).OmitZero().WherePK().Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
