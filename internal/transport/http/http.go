package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to comments services
type Handler struct {
	Router *mux.Router
}

// NewHandler - returns a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - Set up all routes for the application
func (h *Handler) SetupRoutes() {
	fmt.Println("Seting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "I am alive")
	})
}
