package service

import (
	"context"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
)

func (s *Service) ValidateUser(ctx context.Context, email string, pass string) (*types.User, bool, error) {
	return nil, true, nil
}
