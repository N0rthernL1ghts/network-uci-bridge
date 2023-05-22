package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
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
		logger.Error(fmt.Sprintf("Error connecting to server: %v", err))
		os.Exit(1)
	}
	defer client.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go client.Listen(&wg)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		client.Send(input)
	}

	wg.Wait()
}
