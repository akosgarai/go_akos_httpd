package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/akosgarai/go_akos_httpd/httpd"
	"github.com/akosgarai/go_akos_httpd/store"
)

// DefaultHTTPAddr is the default HTTP bind address.
const DefaultHTTPAddr = ":8080"

// Parameters
var httpAddr string

// init initializes this package.
func init() {
	flag.StringVar(&httpAddr, "addr", DefaultHTTPAddr, "Set the HTTP bind address")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
}

// main is the entry point for the service.
func main() {
	flag.Parse()

	s := store.New()
	if err := s.Open(); err != nil {
		log.Fatalf("failed to open store: %s", err.Error())
	}

	h := httpd.New(httpAddr, s)
	if err := h.Start(); err != nil {
		log.Fatalf("failed to start HTTP service: %s", err.Error())
	}

	log.Println("httpd started successfully")

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Block until one of the signals above is received
	select {
	case <-signalCh:
		log.Println("signal received, shutting down...")
		if err := s.Close(); err != nil {
			log.Println("failed to close store:", err.Error())
		}
		if err := h.Close(); err != nil {
			log.Println("failed to close HTTP service:", err.Error())
		}
	}
}
