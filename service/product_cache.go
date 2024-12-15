package service

import (
	"context"
	"encoding/json"
	"time"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// ProductCache struct for managing product cache
type ProductCache struct {
	Client *redis.Client
}

// NewProductCache initializes the Redis client
func NewProductCache(redisAddr string) *ProductCache {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr, // Redis address
	})
	return &ProductCache{Client: client}
}

// GetProduct fetches a product from the cache
func (pc *ProductCache) GetProduct(productID int) (string, error) {
	return pc.Client.Get(ctx, cacheKey(productID)).Result()
}

// SetProduct saves a product in the cache
func (pc *ProductCache) SetProduct(productID int, productData interface{}) error {
	data, err := json.Marshal(productData)
	if err != nil {
		return err
	}
	return pc.Client.Set(ctx, cacheKey(productID), data, 10*time.Minute).Err()
}

// InvalidateProduct removes a product from the cache
func (pc *ProductCache) InvalidateProduct(productID int) error {
	return pc.Client.Del(ctx, cacheKey(productID)).Err()
}

func cacheKey(productID int) string {
	return fmt.Sprintf("product:%d", productID)
}
