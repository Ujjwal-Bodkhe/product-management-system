package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

// Product struct represents the product table in the database.
type Product struct {
	ID                  int       `json:"id"`
	UserID              int       `json:"user_id"`
	ProductName         string    `json:"product_name"`
	ProductDescription  string    `json:"product_description"`
	ProductImages       []string  `json:"product_images"`
	CompressedImages    []string  `json:"compressed_images"`
	ProductPrice        float64   `json:"product_price"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

// CreateProduct inserts a new product into the database.
func CreateProduct(db *sql.DB, product *Product) (int, error) {
	query := `INSERT INTO products (user_id, product_name, product_description, product_price, product_images, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id`

	var productID int
	err := db.QueryRow(query, product.UserID, product.ProductName, product.ProductDescription, product.ProductPrice, pq.Array(product.ProductImages)).Scan(&productID)
	if err != nil {
		return 0, err
	}
	return productID, nil
}



// GetProductByID retrieves a product by its ID.
func GetProductByID(db *sql.DB, id int) (*Product, error) {
	query := `SELECT id, user_id, product_name, product_description, product_price, product_images, compressed_product_images, created_at, updated_at
              FROM products WHERE id = $1`

	row := db.QueryRow(query, id)

	var product Product
	var productImages, compressedImages pq.StringArray
	err := row.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductPrice, &productImages, &compressedImages, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	product.ProductImages = productImages
	product.CompressedImages = compressedImages
	return &product, nil
}

// ListProductsByUser retrieves all products for a specific user.
func ListProductsByUser(db *sql.DB, userID int) ([]Product, error) {
	query := `SELECT id, user_id, product_name, product_description, product_price, product_images, compressed_product_images, created_at, updated_at
              FROM products WHERE user_id = $1`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		var productImages, compressedImages pq.StringArray
		err := rows.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductPrice, &productImages, &compressedImages, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		product.ProductImages = productImages
		product.CompressedImages = compressedImages
		products = append(products, product)
	}
	return products, nil
}
