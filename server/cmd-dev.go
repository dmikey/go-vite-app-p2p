//go:build dev && !app
// +build dev,!app

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)
var headless bool
var port int

func init() {
	// Define the headless flag
	flag.BoolVar(&headless, "headless", false, "Run in headless mode without opening the browser")
	flag.IntVar(&port, "port", 0, "Run in headless mode without opening the browser")
	flag.Parse()
}

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

    listenAddr := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen on a port: %v", err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	serverURL := fmt.Sprintf("http://localhost:%d", port)
	fmt.Printf("Development server listening on %s\n", serverURL)
    if err := http.Serve(listener, nil); err != nil {
        fmt.Println(err)
    }
}
