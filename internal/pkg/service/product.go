package service

import (
	"context"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
)

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

func (s *Service) GetProductInfo(ctx context.Context, id string) (*types.Product, error) {
	model := new(types.ProductModel)
	if err := s.db.NewSelect().Model(model).Where("p.id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}

	return model.ToProduct(), nil
}
