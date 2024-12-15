package service

import (
	"errors"
	"github.com/yourusername/product-management-system/storage"
	"github.com/yourusername/product-management-system/cache"
	"github.com/yourusername/product-management-system/queue"
)

type ProductService struct {
	db           *storage.DB
	redisClient  *cache.RedisClient
	messageQueue *queue.MessageQueue
}

func NewProductService(db *storage.DB, redisClient *cache.RedisClient, messageQueue *queue.MessageQueue) *ProductService {
	return &ProductService{db, redisClient, messageQueue}
}

func (s *ProductService) CreateProduct(product *Product) error {
	// Save product to DB
	err := s.db.SaveProduct(product)
	if err != nil {
		return err
	}

	// Push image URLs to the processing queue
	s.messageQueue.PushImageURLs(product.ProductImages)

	return nil
}

func (s *ProductService) GetProductByID(id string) (*Product, error) {
	// Check cache first
	product, found := s.redisClient.Get(id)
	if found {
		return product, nil
	}

	// If not in cache, fetch from DB
	product, err := s.db.GetProductByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}

	// Cache the product
	s.redisClient.Set(id, product)

	return product, nil
}

func (s *ProductService) GetProductsByUser(userID string) ([]Product, error) {
	// Fetch products for a specific user
	products, err := s.db.GetProductsByUser(userID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
