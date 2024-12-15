package storage

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func InitDB() (*DB, error) {
	connStr := "postgres://username:password@localhost/dbname?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DB{conn: db}, nil
}

func (db *DB) SaveProduct(product *Product) error {
	// Database logic to save product
	query := "INSERT INTO products (user_id, product_name, product_description, product_images, product_price) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.conn.Exec(query, product.UserID, product.ProductName, product.ProductDescription, product.ProductImages, product.ProductPrice)
	return err
}

func (db *DB) GetProductByID(id string) (*Product, error) {
	// Database logic to get product by ID
	query := "SELECT * FROM products WHERE id = $1"
	row := db.conn.QueryRow(query, id)
	var product Product
	err := row.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.ProductPrice)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (db *DB) GetProductsByUser(userID string) ([]Product, error) {
	// Database logic to fetch products by user
	query := "SELECT * FROM products WHERE user_id = $1"
	rows, err := db.conn.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.ProductPrice)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products, nil
}
