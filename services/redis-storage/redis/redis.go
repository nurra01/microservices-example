package redis

import (
	"fmt"
	"os"
	"services/redis-storage/models"
	"services/redis-storage/utils"
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
	ok, err := client.SetNX(verifyID, user, 10*time.Minute).Result() // expiration time is 10 minutes
	if err != nil || !ok {
		return fmt.Errorf("failed to set key with user value. %s", err.Error())
	}

	return nil
}

// SaveToken saves a token in Redis storage for 15 minutes
func SaveToken(msg []byte) error {
	var user *models.RegisterUser
	// convert byte stream to object
	err := utils.FromByteToObject(msg, user)
	if err != nil {
		return fmt.Errorf("failed to deserialize message from kafka. %s", err.Error())
	}
	// create a token from user object
	token, err := utils.CreateToken(user)
	if err != nil {
		return err
	}
	// set a email key with token value
	ok, err := client.SetNX(user.Email, token, 15*time.Minute).Result() // expiration time is 15 minutes
	if err != nil || !ok {
		return fmt.Errorf("failed to set key with token value %s", err.Error())
	}

	return nil
}
