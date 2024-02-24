package main

import (
	"encoding/json"
	"net/http"
)

// apiHandler is an example API handler function.
func apiHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello from the API"}
	json.NewEncoder(w).Encode(response)
}

// RegisterAPIRoutes sets up the API routes.
func RegisterAPIRoutes() {
	http.HandleFunc("/api", apiHandler)
	// Add more API routes here.
}
