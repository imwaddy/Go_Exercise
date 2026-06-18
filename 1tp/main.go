package main

import (
	"context"

	json "encoding/json/v2"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Set up a Redis client options
	options := &redis.Options{
		Addr:     "localhost:8001", // Replace with your Redis server address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	}

	json.Encoder

	// Create a Redis client
	client := redis.NewClient(options)

	// Example: Set a key-value pair
	key := "exampleKey"
	value := "exampleValue"

	err := client.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
		return
	}
	fmt.Printf("Set key '%s' with value '%s'\n", key, value)

	// Example: Get the value for a key
	retrievedValue, err := client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		fmt.Printf("Key '%s' does not exist\n", key)
	} else if err != nil {
		fmt.Println("Error getting key:", err)
	} else {
		fmt.Printf("Retrieved value for key '%s': '%s'\n", key, retrievedValue)
	}

	// Close the Redis client connection
	err = client.Close()
	if err != nil {
		fmt.Println("Error closing Redis connection:", err)
	}
}
