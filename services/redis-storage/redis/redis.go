package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

// pointer to global redis client
var client *redis.Client

// ConnectRedis sets up connection with redis
func ConnectRedis() error {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

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

// SaveUser saves user with unique key in redis with expiration time
func SaveUser(verifyID string, user []byte) error {
	ok, err := client.SetNX(verifyID, user, 10*time.Minute).Result()
	if err != nil || !ok {
		return fmt.Errorf("failed to set key with user value. %s", err.Error())
	}

	return nil
}
