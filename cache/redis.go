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
		Addr: "localhost:6379", // Adjust for your Redis server
	})
	return &RedisClient{client: client}
}

// Set stores product data in Redis
func (r *RedisClient) Set(id string, product *models.Product) error {
	// Serialize the product struct to JSON
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}

	// Store the serialized product in Redis with the provided id
	return r.client.Set(context.Background(), id, data, 0).Err()
}

// Get retrieves product data from Redis
func (r *RedisClient) Get(id string) (interface{}, bool) {
	// Retrieve the serialized product data from Redis
	data, err := r.client.Get(context.Background(), id).Result()
	if err != nil {
		return nil, false
	}

	// Deserialize the data back into a product struct
	var product models.Product
	err = json.Unmarshal([]byte(data), &product)
	if err != nil {
		return nil, false
	}

	return &product, true
}
