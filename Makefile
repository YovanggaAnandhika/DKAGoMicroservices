# Set the root project path and release folder path
ROOT_DIR := $(shell pwd)
RELEASE_DIR := $(ROOT_DIR)/release

# Default target
all: build-core-account

# Build All Proto
build-proto:
	@echo "Compile All Proto"
	protoc --go_out=./generated --go-grpc_out=./generated api/**/*.proto

# Build core-account
build-core-account: clean build-proto
	@echo "Creating release directory if it doesn't exist..."
	mkdir -p $(RELEASE_DIR)
	@echo "Building core-account..."
	cd cmd/services/core/account && go build -o $(RELEASE_DIR)/services/core/account

# Run core-account
run-core-account: build-core-account
	@echo "Running core-account..."
	$(RELEASE_DIR)/services/core/account

# Clean built binaries
clean:
	@echo "Cleaning up..."
	rm -rf $(RELEASE_DIR)/*
	rm -rf generated/*