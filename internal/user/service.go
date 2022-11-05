package user

import (
	"context"
	"fmt"
)

type (
	service struct {
		cache userService
	}

	validationError struct {
		error
	}
)

func newService(cache *cache) *service {
	return &service{cache: cache}
}

func (s *service) Create(ctx context.Context, user User) (int64, error) {
	if err := user.Validate(); err != nil {
		return 0, fmt.Errorf("validate user: %w", err)
	}
	return s.cache.Create(ctx, user)
}

func (s *service) Get(ctx context.Context, userID int64) (User, error) {
	if userID <= 0 {
		return User{}, validationError{errInvalidID}
	}
	return s.cache.Get(ctx, userID)
}

func (s *service) Update(ctx context.Context, user User) error {
	if err := user.Validate(); err != nil {
		return fmt.Errorf("validate user: %w", err)
	}
	return s.cache.Update(ctx, user)
}

func (s *service) Delete(ctx context.Context, userID int64) error {
	if userID <= 0 {
		return validationError{errInvalidID}
	}
	return s.cache.Delete(ctx, userID)
}
