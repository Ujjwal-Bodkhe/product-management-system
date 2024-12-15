package models

import (
	"database/sql"
	"time"
)

// User struct represents the user table in the database.
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUser inserts a new user into the database.
func CreateUser(db *sql.DB, user *User) (int, error) {
	query := `INSERT INTO users (username, email, password, created_at, updated_at) 
              VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`

	var userID int
	err := db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// GetUserByID retrieves a user by their ID.
func GetUserByID(db *sql.DB, id int) (*User, error) {
	query := `SELECT id, username, email, created_at, updated_at FROM users WHERE id = $1`

	row := db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail retrieves a user by their email.
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	query := `SELECT id, username, email, password, created_at, updated_at FROM users WHERE email = $1`

	row := db.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
