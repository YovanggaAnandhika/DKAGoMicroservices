# Set the root project path and release folder path
ROOT_DIR := $(shell pwd)
RELEASE_DIR := $(ROOT_DIR)/release

# Default target
all: build-core-account

# Build All Proto
proto:
	@echo "Compile All Proto"
	find api/ -type f -name "*.proto" | xargs protoc --go_out=./generated --go-grpc_out=./generated

inject: proto
	@echo "inject proto"
	find generated/ -type f -name "*.pb.go" | xargs -I {} protoc-go-inject-tag --input={}

# run dev
dev-parking-config:
	@echo "Building parking-config..."
	cd cmd/services/parking/config && go run main.go

# Build core-account
build-core-account: clean proto
	@echo "Creating release directory if it doesn't exist..."
	mkdir -p $(RELEASE_DIR)
	@echo "Building core-account..."
	cd cmd/services/core/account && go build -o $(RELEASE_DIR)/services/core/account

# run dev
dev-core-account:
	@echo "Building core-account..."
	cd cmd/services/core/account && go run main.go

# Run core-account
run-core-account: build-core-account
	@echo "Running core-account..."
	$(RELEASE_DIR)/services/core/account

# Clean built binaries
clean:
	@echo "Cleaning up..."
	rm -rf $(RELEASE_DIR)/*
	rm -rf generated/*