package users_service

import (
	"context"
	"fmt"

	"github.com/jettyjunk/goland-todoapp/internal/core/domain"
	core_errors "github.com/jettyjunk/goland-todoapp/internal/core/errors"
)

func (s *UsersService) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be non-negativ: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be non-negativ: %w", core_errors.ErrInvalidArgument)
	}

	users, err := s.usersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return []domain.User{}, fmt.Errorf("get users from repository: %w", err)
	}

	return users, nil
}
