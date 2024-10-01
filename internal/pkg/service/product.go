package service

import (
	"context"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
	"github.com/google/uuid"
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
