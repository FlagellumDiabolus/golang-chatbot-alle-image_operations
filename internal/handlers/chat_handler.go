package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang-chatbot-alle-image_operations/internal/database"
)

const rasaServerURL = "http://localhost:5005/model/parse"

// ChatHandler handles chat messages
// ChatHandler handles chat messages
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON request body
	var req struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// If message is not related to file operations, send it to Rasa server for processing
	if !isFileOperation(req.Message) {
		// Create request payload for Rasa server
		payload := map[string]string{
			"text": req.Message,
		}
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			http.Error(w, "Failed to marshal request payload", http.StatusInternalServerError)
			return
		}

		// Send request to Rasa server
		resp, err := http.Post(rasaServerURL, "application/json", bytes.NewReader(payloadBytes))
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send request to Rasa server: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Check response status code
		if resp.StatusCode != http.StatusOK {
			http.Error(w, fmt.Sprintf("Rasa server returned non-OK status code: %d", resp.StatusCode), http.StatusInternalServerError)
			return
		}

		// Decode Rasa server response
		var rasaResp struct {
			Intent string `json:"intent"`
			// Add other fields as needed (e.g., entities)
		}
		if err := json.NewDecoder(resp.Body).Decode(&rasaResp); err != nil {
			http.Error(w, "Failed to decode Rasa server response", http.StatusInternalServerError)
			return
		}

		// Handle Rasa server response
		// Add your logic here to process the Rasa server response
		// For demonstration purposes, we just echo back the intent
		response := fmt.Sprintf("Intent: %s", rasaResp.Intent)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"response": response})
	} else {
		// Process file operation
		processFileOperation(w, req.Message)
	}
}

// isFileOperation checks if the message is related to file operations
func isFileOperation(message string) bool {
	return strings.Contains(strings.ToLower(message), "save") ||
		strings.Contains(strings.ToLower(message), "retrieve") ||
		strings.Contains(strings.ToLower(message), "get")
}

// processFileOperation processes recognized file operations
func processFileOperation(w http.ResponseWriter, message string) {
	switch {
	case strings.Contains(strings.ToLower(message), "save"):
		saveImage(w, message)
	case strings.Contains(strings.ToLower(message), "retrieve") || strings.Contains(strings.ToLower(message), "get"):
		retrieveImage(w, message)
	default:
		http.Error(w, "File operation not supported", http.StatusBadRequest)
	}
}

// saveImage handles the save image operation
func saveImage(w http.ResponseWriter, message string) {
	// Parse image name from message
	imageName := extractImageName(message)
	if imageName == "" {
		http.Error(w, "Image name not provided", http.StatusBadRequest)
		return
	}

	// Parse image URL from message
	imageURL := extractImageURL(message)
	if imageURL == "" {
		http.Error(w, "Image URL not provided", http.StatusBadRequest)
		return
	}

	// Save image details to database
	err := database.SaveImage(imageName, imageURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save image: %v", err), http.StatusInternalServerError)
		return
	}

	response := "Image saved successfully."
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"response": response})
}

// retrieveImage handles the retrieve image operation
func retrieveImage(w http.ResponseWriter, message string) {
	// Parse image name from message
	imageName := extractImageName(message)
	if imageName == "" {
		http.Error(w, "Image name not provided", http.StatusBadRequest)
		return
	}

	// Retrieve image URL from database
	imageURL, err := database.RetrieveImage(imageName)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, fmt.Sprintf("Image '%s' not found", imageName), http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to retrieve image: %v", err), http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("Retrieved image link: %s", imageURL)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"response": response})
}

// extractImageName extracts the image name from the message
func extractImageName(message string) string {
	parts := strings.Fields(message)
	if len(parts) < 3 {
		return ""
	}
	return parts[2]
}

// extractImageURL extracts the image URL from the message
func extractImageURL(message string) string {
	parts := strings.Fields(message)
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}
