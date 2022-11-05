package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

func newRedis(cfg Config) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Username: cfg.User,
		Password: cfg.Pass,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return cli, cli.Ping(ctx).Err()
}
