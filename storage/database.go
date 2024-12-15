package storage

import (
	"database/sql"
	"log"
	"github.com/Ujjwal-Bodkhe/product-management-system/models"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func InitDB() *DB {
	// Connect to PostgreSQL database
	connStr := "user=postgres dbname=product_manager sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}

func (db *DB) SaveProduct(product *models.Product) error {
	_, err := db.Exec(
		"INSERT INTO products (user_id, product_name, product_description, product_images, product_price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NOW(), NOW())",
		product.UserID, product.ProductName, product.ProductDescription, product.ProductImages, product.ProductPrice,
	)
	return err
}

func (db *DB) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	err := db.QueryRow(
		"SELECT id, user_id, product_name, product_description, product_images, product_price, created_at, updated_at FROM products WHERE id = $1",
		id,
	).Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.ProductPrice, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (db *DB) GetProductsByUser(userID string) ([]models.Product, error) {
	rows, err := db.Query(
		"SELECT id, user_id, product_name, product_description, product_images, product_price, created_at, updated_at FROM products WHERE user_id = $1",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.ProductPrice, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
