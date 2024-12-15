package cache

import (
	"context"
	"encoding/json"
	"github.com/Ujjwal-Bodkhe/product-management-system/models"
	"github.com/go-redis/redis/v8"
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

func (r *RedisClient) Set(id string, product *models.Product) error {
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}
	return r.client.Set(context.Background(), id, data, 0).Err()
}

func (r *RedisClient) Get(id string) (interface{}, bool) {
	data, err := r.client.Get(context.Background(), id).Result()
	if err != nil {
		return nil, false
	}

	var product models.Product
	err = json.Unmarshal([]byte(data), &product)
	if err != nil {
		return nil, false
	}

	return &product, true
}
