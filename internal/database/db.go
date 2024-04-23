package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitializeDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to the database")
	return nil
}

func SaveImage(name, url string) error {
	query := "INSERT INTO images (name, url) VALUES (?, ?)"
	_, err := db.Exec(query, name, url)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}
	return nil
}

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

func ListImages() ([]string, error) {
	query := "SELECT name FROM images"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var images []string

	for rows.Next() {
		var imageName string
		if err := rows.Scan(&imageName); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		images = append(images, imageName)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration: %v", err)
	}

	return images, nil
}
