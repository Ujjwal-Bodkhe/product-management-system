package storage

import (
	"github.com/Ujjwal-Bodkhe/product-management-system/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB() *DB {
	dsn := "user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	return &DB{db}
}

func (db *DB) SaveProduct(product *models.Product) error {
	return db.Create(product).Error
}

func (db *DB) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	err := db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (db *DB) GetProductsByUser(userID string) ([]models.Product, error) {
	var products []models.Product
	err := db.Where("user_id = ?", userID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
