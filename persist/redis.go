package persist

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// RedisClient - stores the connection to Redis server.
var RedisClient *redis.Client

// InitRedisClient - initializes Redis storage driver.
func InitRedisClient(host, port, password string) *redis.Client {
	if RedisClient != nil {
		return RedisClient
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	RedisClient = conn

	return RedisClient
}
