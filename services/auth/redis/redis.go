package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

// pointer to global redis client
var client *redis.Client

// ConnectRedis sets up connection with redis
func ConnectRedis() error {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	// if env variables missing, use default
	if host == "" || port == "" {
		host = "redis"
		port = "6379"
	}

	c := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := c.Ping().Result()
	if err != nil {
		return fmt.Errorf("failed to connect redis. %s", err.Error())
	}
	client = c // set pointer to global vriable
	return nil
}
