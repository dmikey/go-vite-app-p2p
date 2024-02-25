//go:build !dev && !app
// +build !dev,!app

package main

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed assets/*
var embeddedFiles embed.FS
var headless bool
var port int
var logger zerolog.Logger

func init() {
	// Define the headless flag
	flag.BoolVar(&headless, "headless", false, "Run in headless mode without opening the browser")
	flag.IntVar(&port, "port", 0, "Run in headless mode without opening the browser")
	flag.Parse()

	logger = zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	if(!headless) {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

func main() {
	cfg := parseFlags()
	// Signal catching for clean shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	done := make(chan struct{})
	failed := make(chan struct{})

	logger.Info().Bool("headless mode", !headless).Msg("the server is starting")

	assets, err := fs.Sub(embeddedFiles, "assets")
	if err != nil {
		logger.Info().Err(err).Msg("Failed to locate embedded assets")
		return
	}

	// Get the port the server is listening on.
	// Listen on a random port.
	listenHost := "localhost"
	if(headless) {
		listenHost = "0.0.0.0"
	}

	listenAddr := fmt.Sprintf("%s:%d", listenHost, port)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to listen on a port: %v")
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	serverURL := fmt.Sprintf("http://localhost:%d", port)
	logger.Info().Str("serverUrl", serverURL).Msgf("Production server listening on %s", serverURL)

	// Use the http.FileServer to serve the embedded assets.
	http.Handle("/", http.FileServer(http.FS(assets)))

    // Register API routes.
	RegisterAPIRoutes()

	// Create the main context for p2p
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Open the pebble peer database.
	pdb, err := pebble.Open(cfg.PeerDatabasePath, &pebble.Options{Logger: &pebbleNoopLogger{}})
	if err != nil {
		log.Error().Err(err).Str("db", cfg.PeerDatabasePath).Msg("could not open pebble peer database")
	}
	defer pdb.Close()

	// Open the pebble function database.
	fdb, err := pebble.Open(cfg.FunctionDatabasePath, &pebble.Options{Logger: &pebbleNoopLogger{}})
	if err != nil {
		log.Error().Err(err).Str("db", cfg.FunctionDatabasePath).Msg("could not open pebble function database")
	}
	defer fdb.Close()

	// Boot P2P Network
	runP2P(ctx, logger, *cfg, done, failed, pdb, fdb)

	if !headless {
		logger.Info().Msg("Opening browser")
		go func() {
			waitForServer(serverURL)
			openbrowser(serverURL)
		}()
	}

	// Start API in a separate goroutine.
	go func() {
		logger.Info().Str("port", cfg.API).Msg("Node API starting")
		err := http.Serve(listener, nil)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Warn().Err(err).Msg("Closed Server")
			close(failed)
		}
	}()

	select {
		case <-sig:
			logger.Info().Msg("Blockless AVS stopping")
		case <-done:
			logger.Info().Msg("Blockless AVS done")
		case <-failed:
			logger.Info().Msg("Blockless AVS aborted")
	}

	// If we receive a second interrupt signal, exit immediately.
	go func() {
		<-sig
		logger.Warn().Msg("forcing exit")
		os.Exit(1)
	}()
}

func waitForServer(url string) {
	for {
		// Attempt to connect to the server.
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close() // Don't forget to close the response body.
			logger.Info().Msg("App is Running. CTRL+C to quit.")
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
		logger.Fatal().Err(err)
	}

}