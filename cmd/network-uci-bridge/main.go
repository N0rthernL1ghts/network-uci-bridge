package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := NewLogger()

	// Get the TCP server host and port from environment variables
	host := os.Getenv("UCI_TCP_HOST")
	port := os.Getenv("UCI_TCP_PORT")

	// Check if the environment variables are set
	if host == "" || port == "" {
		logger.Fatal("UCI_TCP_HOST or UCI_TCP_PORT environment variables are not set")
	}

	// Initiate connection to the TCP server
	client, err := NewClient(host+":"+port, logger)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error connecting to the server: %v", err))
	}
	defer client.Close()

	// Create a context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run client.Listen with the context
	done := make(chan struct{})
	go client.Listen(ctx, done)

	// Run a separate Goroutine to handle user input
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()

			// If input is "quit", send quit command to the server, cleanup and exit
			if input == "quit" {
				client.Send(input)
				cancel()
				return
			}

			client.Send(input)
		}
	}()

	// Create a channel to signal when the program is exiting
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGTERM)

	// Wait for the program to be terminated or the done channel to be closed
	select {
	case <-exitChan:
		client.Send("quit")
	case <-done:
	}
	<-done
}
