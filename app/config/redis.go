package config

import (
	"context"

	"github.com/Uttamnath64/quixzap/app/storage"
)

func InitRedis(ctx context.Context, addr, password string, db int) (*storage.RedisClient, error) {
	return storage.NewRedisClient(ctx, addr, password, db)
}
