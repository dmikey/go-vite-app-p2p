//go:build dev
// +build dev

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve files directly from the filesystem in development.
	http.Handle("/", http.FileServer(http.Dir("assets/")))

    // Register API routes.
	RegisterAPIRoutes()

	fmt.Println("Development server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
