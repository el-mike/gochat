package mocks

import (
	"context"
	"time"

	"github.com/el-Mike/gochat/persist"
	"github.com/stretchr/testify/mock"
)

// RedisCacheMock - basic, reusable mock for Redis Cache client.
type RedisCacheMock struct {
	mock.Mock
}

func NewRedisCacheMock() *RedisCacheMock {
	return &RedisCacheMock{}
}

// Get - Get method mock implementation.
func (rc *RedisCacheMock) Get(
	ctx context.Context,
	key string,
) *persist.CacheResponse {
	args := rc.Called(ctx, key)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*persist.CacheResponse)
}

// Set - Set method mock implementation.
func (rc *RedisCacheMock) Set(
	ctx context.Context,
	key string,
	value interface{},
	expiration time.Duration,
) *persist.CacheResponse {
	args := rc.Called(ctx, key, value, expiration)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*persist.CacheResponse)
}

// Del - Del method mock implementation.
func (rc *RedisCacheMock) Del(ctx context.Context, keys ...string) *persist.CacheResponse {
	args := rc.Called(ctx, keys)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*persist.CacheResponse)
}

// GetDefaultCacheResponse - returns default, empty CacheResponse.
func GetDefaultCacheResponse() *persist.CacheResponse {
	return persist.NewCacheResponse()
}

// GetErrorCacheResponse - returns CacheResponse with given error.
func GetErrorCacheResponse(err error) *persist.CacheResponse {
	res := persist.NewCacheResponse()
	res.SetErr(err)

	return res
}
