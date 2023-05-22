package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Client struct {
	conn   net.Conn
	logger *Logger
}

func NewClient(address string, logger *Logger) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn:   conn,
		logger: logger,
	}

	return client, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Send(message string) {
	_, err := c.conn.Write([]byte(message + "\n"))
	if err != nil {
		c.logger.Error(fmt.Sprintf("Error sending message: %v", err))
	}
}

func (c *Client) Listen(wg *sync.WaitGroup) {
	defer wg.Done()

	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println(message)
	}
}
