package repositories

import (
	"database/sql"
)

// User represents the structure of our resource
type User struct {
	ID    int
	Name  string
	Email string
}

// CreateUser inserts a new user into the database
func CreateUser(db *sql.DB, user User) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO users(name, email) VALUES($1, $2) RETURNING id", user.Name, user.Email).Scan(&id)
	return id, err
}

// GetUser retrieves a user from the database by ID
func GetUser(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

// UpdateUser updates a user's information in the database
func UpdateUser(db *sql.DB, user User) error {
	_, err := db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, user.ID)
	return err
}

// DeleteUser removes a user from the database by ID
func DeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
