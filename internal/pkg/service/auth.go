package service

import (
	"context"

	"github.com/ak1m1tsu/jokerge/internal/pkg/types"
)

// ValidateUser проверяет наличие пользователя по email и сравнивает переданный пароль с найденым.
func (s *Service) ValidateUser(ctx context.Context, email string, pass string) (*types.User, bool, error) {
	user := new(types.User)

	if err := s.db.NewSelect().Model(user).Where("email = ?", email).Scan(ctx); err != nil {
		return nil, false, err
	}

	if user.Password != pass {
		return nil, false, nil
	}

	return user, true, nil
}
