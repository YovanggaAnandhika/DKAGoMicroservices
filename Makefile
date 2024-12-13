# Default target
all: build-core-account

# Build All Proto
build-proto:
	@echo "Compile All Proto"
	protoc --go_out=./generated --go-grpc_out=./generated api/**/*.proto

# Build core-account
build-core-account: clean build-proto
	@echo "Building core-account..."
	cd cmd/services/core/account && go build -o account

# Run core-account
run-core-account: build-core-account
	@echo "Running core-account..."
	./cmd/services/core/account/account

# Clean built binaries
clean:
	@echo "Cleaning up..."
	rm -f cmd/services/core/account/account
	rm -rf generated/*