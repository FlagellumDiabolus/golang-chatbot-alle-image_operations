package server

import (
	"net/http"

	"golang-chatbot-alle-image_operations/internal/handlers"
)

// Server represents the HTTP server
type Server struct {
	mux *http.ServeMux
}

// NewServer creates a new instance of Server
func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

// SetupRoutes sets up routes for the server
func (s *Server) SetupRoutes() {
	// Handle chat messages
	s.mux.HandleFunc("/chat", handlers.ChatHandler)

	// Handle image operations
	s.mux.HandleFunc("/save-image", handlers.SaveImageHandler)
	s.mux.HandleFunc("/retrieve-image", handlers.RetrieveImageHandler)
}

// ServeHTTP implements the http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Serve HTTP requests using the server's mux
	s.mux.ServeHTTP(w, r)
}
