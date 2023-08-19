package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func createRedisClient() *redis.Client {
	fmt.Println("Redis connection is completed")
	fmt.Println("")
	return redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0,
	})
}
