package persist

import (
	"context"
	"time"
)

// Cache - basic, common cache interface.
type Cache interface {
	// Get - get a value by given key.
	Get(ctx context.Context, key string) *CacheResponse

	// Set - set given key to the passed value.
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *CacheResponse

	// Del - remove value under given key.
	Del(ctx context.Context, keys ...string) *CacheResponse
}

// CacheResponse - basic cache response.
type CacheResponse struct {
	err error
}

// NewCacheResponse - returns CacheResponse instance.
func NewCacheResponse() *CacheResponse {
	return &CacheResponse{}
}

// Err - returns an error that occured during cache operation.
func (cr *CacheResponse) Err() error {
	return cr.err
}

// SetError - sets given error on CacheResponse instance.
func (cr *CacheResponse) SetErr(err error) {
	cr.err = err
}
