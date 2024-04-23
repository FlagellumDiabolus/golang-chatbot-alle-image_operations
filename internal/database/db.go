package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitializeDB initializes the SQLite database connection
func InitializeDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Check database connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to the database")
	return nil
}

// SaveImage saves image details to the SQLite database
func SaveImage(name, url string) error {
	query := "INSERT INTO images (name, url) VALUES (?, ?)"
	_, err := db.Exec(query, name, url)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}
	return nil
}

// RetrieveImage retrieves the URL of an image from the SQLite database based on its name
func RetrieveImage(name string) (string, error) {
	var imageURL string
	query := "SELECT url FROM images WHERE name = ?"
	err := db.QueryRow(query, name).Scan(&imageURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("image not found: %v", err)
		}
		return "", fmt.Errorf("failed to retrieve image: %v", err)
	}
	return imageURL, nil
}

// ListImages retrieves the list of all images from the database
func ListImages() ([]string, error) {
	// Query to select all images
	query := "SELECT name FROM images"

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	// Initialize a slice to store image details
	var images []string

	// Iterate over the rows
	for rows.Next() {
		var imageName string
		// Scan the row into the image struct
		if err := rows.Scan(&imageName); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		// Append the image to the slice
		images = append(images, imageName)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration: %v", err)
	}

	return images, nil
}
