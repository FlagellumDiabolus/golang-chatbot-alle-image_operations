package server

import (
	"net/http"

	"golang-chatbot-alle-image_operations/internal/handlers"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) SetupRoutes() {
	s.mux.HandleFunc("/chat", handlers.ChatHandler)

	s.mux.HandleFunc("/save-image", handlers.SaveImageHandler)
	s.mux.HandleFunc("/retrieve-image", handlers.RetrieveImageHandler)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
