package redis

import (
	"fmt"
	"os"

	"services/user/models"
	"services/user/utils"

	"github.com/go-redis/redis"
)

// pointer to global redis client
var client *redis.Client

// ConnectRedis sets up connection with redis
func ConnectRedis() error {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	// if missing env variables, use default
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

// GetUser returns a user from redis if he/she is still available
func GetUser(userID string) (*models.RegisterUser, error) {
	resp, err := client.Get(userID).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user from redis. %s", err.Error())
	} else {
		user := &models.RegisterUser{}
		err = utils.FromByteToObject([]byte(resp), user)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize user from redis. %s", err.Error())
		}
		return user, nil
	}
}
