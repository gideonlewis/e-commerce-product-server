package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	pkgredis "github.com/gideonlewis/e-commerce-product-server/pkg/redis"
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	client *redis.Client
}

func NewService(addr, password string) *Cache {
	client, err := pkgredis.NewConnection(addr, password)
	if err != nil {
		panic(err)
	}

	return &Cache{client: client}
}

func (c *Cache) Get(key string, value interface{}) error {
	data, err := c.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return fmt.Errorf("cache miss for key %q", key)
	} else if err != nil {
		return fmt.Errorf("failed to get value for key %q: %v", key, err)
	}

	if err := json.Unmarshal([]byte(data), value); err != nil {
		return fmt.Errorf("failed to unmarshal cache value for key %q: %v", key, err)
	}

	return nil
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal cache value for key %q: %v", key, err)
	}

	if err := c.client.Set(context.Background(), key, data, duration).Err(); err != nil {
		return fmt.Errorf("failed to set value for key %q: %v", key, err)
	}

	return nil
}

func (c *Cache) Delete(key string) error {
	if err := c.client.Del(context.Background(), key).Err(); err != nil {
		return fmt.Errorf("failed to delete value for key %q: %v", key, err)
	}
	return nil
}
