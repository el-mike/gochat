package mocks

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/mock"
)

// GetDefaultStatusCmd - returns default, empty StatusCmd.
func GetDefaultStatusCmd() *redis.StatusCmd {
	return &redis.StatusCmd{}
}

// GetDefaultIntCmd - returns default, empty IntCmd.
func GetDefaultIntCmd() *redis.IntCmd {
	return &redis.IntCmd{}
}

// RedisCacheMock - basic mock for Redis Cache client.
type RedisCacheMock struct {
	mock.Mock
}

func NewRedisCacheMock() *RedisCacheMock {
	return &RedisCacheMock{}
}

// Set - Set method mock implementation.
func (rc *RedisCacheMock) Set(
	ctx context.Context,
	key string,
	value interface{},
	expiration time.Duration,
) *redis.StatusCmd {
	return GetDefaultStatusCmd()
}

// Del - Del method mock implementation.
func (rc *RedisCacheMock) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return GetDefaultIntCmd()
}
