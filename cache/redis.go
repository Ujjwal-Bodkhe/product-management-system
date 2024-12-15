package cache

import (
	"github.com/go-redis/redis/v8"
	"context"
)

type RedisClient struct {
	client *redis.Client
}

func InitRedis() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &RedisClient{client: client}
}

func (r *RedisClient) Get(key string) (*Product, bool) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, false
	}
	// Assuming Product is stored in JSON format
	var product Product
	// Unmarshal val to product
	return &product, true
}

func (r *RedisClient) Set(key string, product *Product) {
	// Serialize product to JSON and set it in Redis
}
