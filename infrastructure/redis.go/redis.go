package redisUtil

import (
	"context"
	"log"

	redis "github.com/go-redis/redis/v8"
)

// NewClient creates a new Redis client and returns it.
func NewClient() *redis.Client {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",                // No password set
		DB:       0,                 // Use the default database
	})
	err := Ping(client)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return client
}

// Ping checks if the Redis server is reachable.
func Ping(client *redis.Client) error {
	// Ping the Redis server to check if it's reachable
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	if pong != "PONG" {
		return redis.Nil
	}
	return nil
}
