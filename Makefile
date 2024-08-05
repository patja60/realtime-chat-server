.PHONY: test coverage clean-test build run docker-run clean

# Variables
BINARY_NAME=main
DOCKER_COMPOSE=compose
BUILD_DIR=./cmd/realtime-chat-server

# Default target: run all tests
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...

# Generate coverage report
coverage: test
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html

# Clean up coverage files
clean-test:
	@echo "Cleaning up..."
	rm -f coverage.out coverage.html

# Default target: build the Go application
build:
	@echo "Building the application..."
	go build -o $(BINARY_NAME) $(BUILD_DIR)

# Run the server locally
run: build
	@echo "Running the server..."
	./$(BINARY_NAME)

# Run the server using Docker Compose
docker-run:
	@echo "Running the server with Docker Compose..."
	docker $(DOCKER_COMPOSE) up --build

# Run the server using Docker Compose
docker-run-dependency:
	@echo "Running the dependency with Docker Compose..."
	docker $(DOCKER_COMPOSE) up --build --scale app=0

# Clean up binaries and Docker containers
clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)
	$(DOCKER_COMPOSE) down -v