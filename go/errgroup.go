package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func timeoutAfter(ctx context.Context, name string, timeout time.Duration) func() error {
	return func() error {
		select {
		case <-time.After(timeout):
			err := errors.New(name + ": time out")
			log.Print(err)
			return err
		case <-ctx.Done():
			log.Printf(name + ": context canceled, exiting")
			return nil
		}
	}
}

func listenHTTP(ctx context.Context, port int) (startfn, stopfn func() error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		f := func() error { return err } // fail immediately
		return f, f
	}
	srv := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			const delay = 30 * time.Second
			log.Printf("in http.Server.Handler... (sleeping for %s)", delay)
			time.Sleep(delay) // simulate very slow handler
		}),
	}
	startfn = func() error {
		if err := srv.Serve(l); err != nil {
			log.Printf("http.Server.ListenAndServe(:%d) exited, returned %s", port, err)
			return err
		}
		return nil
	}
	stopfn = func() error {
		shutdown := func() error {
			const timeout = 3 * time.Second
			log.Printf("context canceled; asking webserver to shutdown (with %s timeout)", timeout)
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Printf("http.Server.Shutdown returned an error: %s", err)
				log.Printf("closing listen socket")
				l.Close()
				return err
			}
			log.Print("http.Server.Shutdown returned ok")
			return nil
		}
		<-ctx.Done()
		return shutdown()
	}
	return startfn, stopfn
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Create a group of goroutines.
	g, ctx := errgroup.WithContext(ctx)
	g.Go(timeoutAfter(ctx, "alpha", 60*time.Second))
	g.Go(timeoutAfter(ctx, "bravo", 30*time.Second))
	g.Go(timeoutAfter(ctx, "delta", 90*time.Second))
	{
		f1, f2 := listenHTTP(ctx, 5615)
		g.Go(f1)
		g.Go(f2)
	}

	// Cancel the parent context on signal.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	log.Printf("installed signal handlers for signals %d, %d", os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-c
		log.Printf("received signal %d; canceling parent context", s)
		cancel()
	}()

	// Wait for the group to exit.
	if err := g.Wait(); err != nil {
		log.Printf("g.Wait() returned an error: %s", err)
	}
}
