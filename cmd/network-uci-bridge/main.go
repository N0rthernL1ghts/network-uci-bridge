package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sync"
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

	// Create a channel to signal when the program is exiting
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)
	go client.Listen(&wg)

	// Run a separate Goroutine to handle user input
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			client.Send(input)
		}
	}()

	// Wait for the program to be terminated
	<-exitChan

	// Handle cleanup and send "quit" message
	cleanup(client)
	wg.Wait()
}

func cleanup(client *Client) {
	client.Send("quit")
	client.Close()
}
