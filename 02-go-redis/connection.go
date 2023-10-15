package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func connectCache() *redis.Client {
	// client init
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379", // docker-compose service name
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Error init client " + err.Error())
	}

	return client
}
