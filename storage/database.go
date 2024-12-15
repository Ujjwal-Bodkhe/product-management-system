package storage

import (
	"github.com/Ujjwal-Bodkhe/product-management-system/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func NewDB(dataSourceName string) *DB {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	return &DB{db}
}

func (db *DB) SaveProduct(product *models.Product) error {
	_, err := db.Exec(`INSERT INTO products (user_id, product_name, product_description, product_price, product_images, compressed_images, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())`, 
		product.UserID, product.ProductName, product.ProductDescription, product.ProductPrice, product.ProductImages, product.CompressedImages)
	return err
}

func (db *DB) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	err := db.Get(&product, "SELECT * FROM products WHERE id = $1", id)
	return &product, err
}

func (db *DB) GetProductsByUser(userID string) ([]models.Product, error) {
	var products []models.Product
	err := db.Select(&products, "SELECT * FROM products WHERE user_id = $1", userID)
	return products, err
}
