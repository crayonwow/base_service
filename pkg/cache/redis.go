package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

type (
	redisCache struct {
		prefix string
		ttl    time.Duration
		r      *redis.Client
	}
)

func (r *redisCache) Set(ctx context.Context, c Cachable) error {
	b, err := c.MarshalJSON()
	if err != nil {
		return fmt.Errorf("marhsal: %w", err)
	}
	return r.r.SetNX(ctx, r.key(c.Key()), string(b), r.ttl).Err()
}

func (r *redisCache) Get(ctx context.Context, c Cachable) error {
	result, err := r.r.GetEx(ctx, r.key(c.Key()), r.ttl).Result()
	if err != nil {
		return err
	}
	return c.UnmarshalJSON([]byte(result))
}

func (r *redisCache) Drop(ctx context.Context, k Key) error {
	return r.r.Del(ctx, r.key(k)).Err()
}

func (r *redisCache) Invalidate(_ context.Context) error {
	//TODO implement via lua script
	return nil
}

func NewRedisCache(r *redis.Client, prefix string, ttl int) Cache {
	return &redisCache{
		r:      r,
		prefix: prefix,
		ttl:    time.Second * time.Duration(ttl),
	}
}

func (r *redisCache) key(k Key) string {
	return r.prefix + ":" + k.String()
}
