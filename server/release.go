//go:build !dev
// +build !dev

package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed assets/*
var embeddedFiles embed.FS

func main() {
	assets, err := fs.Sub(embeddedFiles, "assets")
	if err != nil {
		fmt.Println("Failed to locate embedded assets:", err)
		return
	}

	// Use the http.FileServer to serve the embedded assets.
	http.Handle("/", http.FileServer(http.FS(assets)))

    // Register API routes.
	RegisterAPIRoutes()

	fmt.Println("Production server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
