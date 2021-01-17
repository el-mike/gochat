package persist

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

// InitRedisClient - initializes Redis storage driver.
func InitRedisClient(host, port, password string) *redis.Client {
	if redisClient != nil {
		return redisClient
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	redisClient = conn

	return redisClient
}
