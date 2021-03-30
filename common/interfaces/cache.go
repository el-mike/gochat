package interfaces

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache - basic, common Redis interface
type RedisCache interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}
