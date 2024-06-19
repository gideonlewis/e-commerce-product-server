package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewConnection(addr, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // use default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return client, nil
}
