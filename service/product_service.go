package service

import (
	
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
	// Save product to the DB
	err := s.db.SaveProduct(product)
	if err != nil {
		return err
	}

	// Push image URLs to the message queue for processing
	s.messageQueue.PushImageURLs(product.ProductImages)

	return nil
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	// Check cache first
	product, found := s.redisClient.Get(id)
	if found {
		// Type assert the result to *models.Product
		return product.(*models.Product), nil
	}

	// Fetch from DB if not found in cache
	product, err := s.db.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	// Cache the product - ensure product is properly type-asserted
	err = s.redisClient.Set(id, product.(*models.Product))
	if err != nil {
		return nil, err
	}

	return product.(*models.Product), nil
}


func (s *ProductService) GetProductsByUser(userID string) ([]models.Product, error) {
	products, err := s.db.GetProductsByUser(userID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
