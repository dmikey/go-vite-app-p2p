//go:build !dev
// +build !dev

package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

//go:embed assets/*
var embeddedFiles embed.FS
var headless bool
var port int

func init() {
	// Define the headless flag
	flag.BoolVar(&headless, "headless", false, "Run in headless mode without opening the browser")
	flag.IntVar(&port, "port", 0, "Run in headless mode without opening the browser")
	flag.Parse()
}

func main() {
	assets, err := fs.Sub(embeddedFiles, "assets")
	if err != nil {
		fmt.Println("Failed to locate embedded assets:", err)
		return
	}

	// Get the port the server is listening on.
	// Listen on a random port.
	listenAddr := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen on a port: %v", err)
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	serverURL := fmt.Sprintf("http://localhost:%d", port)
	fmt.Printf("Production server listening on %s\n", serverURL)

	// Use the http.FileServer to serve the embedded assets.
	http.Handle("/", http.FileServer(http.FS(assets)))

    // Register API routes.
	RegisterAPIRoutes()

	if !headless {
		fmt.Println("Opening browser")
		go func() {
			waitForServer(serverURL)
			openbrowser(serverURL)
		}()
	}

	if err := http.Serve(listener, nil); err != nil {
		fmt.Println(err)
	}
}

func waitForServer(url string) {
	for {
		// Attempt to connect to the server.
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close() // Don't forget to close the response body.
			fmt.Println("App is Running. CTRL+C to quit.")
			return
		}
		// Close the unsuccessful response body to avoid leaking resources.
		if resp != nil {
			resp.Body.Close()
		}
		// Wait for a second before trying again.
		time.Sleep(1 * time.Second)
	}
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}