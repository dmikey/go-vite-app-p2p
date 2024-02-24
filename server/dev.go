//go:build dev && !app
// +build dev,!app

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
    // Register API routes.
    RegisterAPIRoutes()

    // Setup reverse proxy to Vite server
    viteServerURL, _ := url.Parse("http://localhost:5173")
    proxy := httputil.NewSingleHostReverseProxy(viteServerURL)

    // Handle all other requests by forwarding them to the Vite server
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        proxy.ServeHTTP(w, r)
    })

    fmt.Println("Development server listening on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
