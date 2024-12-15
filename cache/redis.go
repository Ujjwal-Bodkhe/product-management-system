package cache

import (
	"encoding/json"
	"strconv"

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

// Set stores product data in Redis
func (r *RedisClient) Set(id int, product *models.Product) error {
    // Example: Convert product to JSON string or use another serialization format
    data, err := json.Marshal(product)
    if err != nil {
        return err
    }

    return r.client.Set(strconv.Itoa(id), data, 0).Err()
}

// Get retrieves product data from Redis
func (r *RedisClient) Get(id int) (interface{}, bool) {
    // Retrieve product data from Redis
    data, err := r.client.Get(strconv.Itoa(id)).Result()
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

