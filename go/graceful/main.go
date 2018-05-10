package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Parse command line flags.
	port := flag.Int("port", 8080, "listen port")
	flag.Parse()

	// Create an HTTP server.
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", *port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "method=%q host=%q path=%q\n", r.Method, r.URL.Host, r.URL.Path)
		}),
	}

	s := &rpc.Server{...}

	// Gracefully shut down the server if we receive a TERM signal.
	sigs := make(chan os.Signal)
	go func(ch <-chan os.Signal) {
		for s := range ch {
			if s == syscall.SIGTERM {
				s.Shutdown(context.Background())
				return
			}
		}
	}(sigs)
	signal.Notify(sigs, syscall.SIGTERM)

	raven.CapturePanic(func() {...})

	// Start listening for requests.
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "ListenAndServe returned %s\n", err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "server gracefully shut down")
	}
}
