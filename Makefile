.PHONY: build clean

PROJECT_NAME := NetworkUciBridge
BINARY_NAME := uci-bridge
CMD_DIR := ./cmd/network-uci-bridge
BIN_DIR := ./build/release

build:
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(CMD_DIR)

clean:
	rm -rf $(BIN_DIR)

run: build
	$(BIN_DIR)/$(BINARY_NAME)
