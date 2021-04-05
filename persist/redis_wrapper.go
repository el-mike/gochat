package persist

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisWrapper struct {
	redis *redis.Client
}

// RedisCache - Cache service based on Redis.
var RedisCache *redisWrapper

// Get - wrapper for Redis' Get method.
func (rc *redisWrapper) Get(ctx context.Context, key string) *CacheResponse {
	cmd := rc.redis.Get(ctx, key)

	return cacheResponseFromStringCmd(cmd)
}

// Set - wrapper for Redis' Set method.
func (rc *redisWrapper) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *CacheResponse {
	cmd := rc.redis.Set(ctx, key, value, expiration)

	return cacheResponseFromStatusCmd(cmd)
}

// Del - wrapper for Redis' Del method.
func (rc *redisWrapper) Del(ctx context.Context, keys ...string) *CacheResponse {
	cmd := rc.redis.Del(ctx, keys...)

	return cacheResponseFromIntCmd(cmd)
}

// InitRedisClient - initializes Redis storage driver.
func InitRedisCache(host, port, password string) *redisWrapper {
	if RedisCache != nil {
		return RedisCache
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	redisCache := &redisWrapper{
		redis: conn,
	}

	RedisCache = redisCache

	return RedisCache
}
func cacheResponseFromStatusCmd(cmd *redis.StatusCmd) *CacheResponse {
	res := NewCacheResponse()

	if cmd.Err() != nil {
		res.SetErr(cmd.Err())
	}

	return res
}

func cacheResponseFromStringCmd(cmd *redis.StringCmd) *CacheResponse {
	res := NewCacheResponse()

	if cmd.Err() != nil {
		res.SetErr(cmd.Err())
	}

	return res
}

func cacheResponseFromIntCmd(cmd *redis.IntCmd) *CacheResponse {
	res := NewCacheResponse()

	if cmd.Err() != nil {
		res.SetErr(cmd.Err())
	}

	return res
}
