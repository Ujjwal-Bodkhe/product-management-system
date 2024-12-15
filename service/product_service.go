package service

import (
	"errors"
	"github.com/Ujjwal-Bodkhe/product-management-system/models"
	"github.com/Ujjwal-Bodkhe/product-management-system/storage"
	"github.com/Ujjwal-Bodkhe/product-management-system/cache"
	"github.com/Ujjwal-Bodkhe/product-management-system/queue"
)

type ProductService struct {
	db           *storage.DB
	redisClient  *cache.RedisClient
	messageQueue *queue.MessageQueue
}

func NewProductService(db *storage.DB, redisClient *cache.RedisClient, messageQueue *queue.MessageQueue) *ProductService {
	return &ProductService{db, redisClient, messageQueue}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	err := s.db.SaveProduct(product)
	if err != nil {
		return err
	}

	// Push image URLs to the queue for processing
	s.messageQueue.PushImageURLs(product.ProductImages)

	return nil
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	// Check cache first
	product, found := s.redisClient.Get(id)
	if found {
		return product.(*models.Product), nil
	}

	// Fetch from DB if not found in cache
	product, err := s.db.GetProductByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}

	// Cache the product
	s.redisClient.Set(id, product)

	return product, nil
}

func (s *ProductService) GetProductsByUser(userID string) ([]models.Product, error) {
	products, err := s.db.GetProductsByUser(userID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
