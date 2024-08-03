.PHONY: test coverage clean

# Default target: run all tests
test:
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...

# Generate coverage report
coverage: test
	@echo "Generating coverage report..."
	go tool cover -html=coverage.out -o coverage.html

# Clean up coverage files
clean:
	@echo "Cleaning up..."
	rm -f coverage.out coverage.html
