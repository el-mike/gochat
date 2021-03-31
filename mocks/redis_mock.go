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

func GetErrorStatusCmd(err error) *redis.StatusCmd {
	cmd := &redis.StatusCmd{}
	cmd.SetErr(err)

	return cmd
}

// GetDefaultIntCmd - returns default, empty IntCmd.
func GetDefaultIntCmd() *redis.IntCmd {
	return &redis.IntCmd{}
}

func GetErrorIntCmd(err error) *redis.IntCmd {
	cmd := &redis.IntCmd{}
	cmd.SetErr(err)

	return cmd
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
	args := rc.Called(ctx, key, value, expiration)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*redis.StatusCmd)
}

// Del - Del method mock implementation.
func (rc *RedisCacheMock) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	args := rc.Called(ctx, keys)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*redis.IntCmd)
}
