package user

import (
	"context"
	"fmt"

	cachePkg "github.com/crayonwow/base_service/pkg/cache"
	"github.com/go-redis/redis/v9"
)

type (
	cache struct {
		c    cachePkg.Cache
		repo userService
	}
)

func newCache(r *redis.Client, repo *repo) *cache {
	return &cache{
		c:    cachePkg.NewRedisCache(r, "user", 120),
		repo: repo,
	}
}

func (c *cache) Get(ctx context.Context, userID int64) (User, error) {
	u := User{ID: userID}
	err := c.c.Get(ctx, &u)
	if err == nil {
		return u, nil
	}

	return c.repo.Get(ctx, userID)
}

func (c *cache) Update(ctx context.Context, user User) error {
	err := c.repo.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("repo update: %w", err)
	}
	return c.c.Drop(ctx, newUserCacheKey(user.ID))
}

func (c *cache) Delete(ctx context.Context, userID int64) error {
	err := c.repo.Delete(ctx, userID)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return c.c.Drop(ctx, newUserCacheKey(userID))
}

func (c *cache) Create(ctx context.Context, user User) (int64, error) {
	return c.repo.Create(ctx, user)
}
