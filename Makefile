build:
	@echo "Building..."
	@go build ./cmd/cli/main.go godo

run:
	@go run ./cmd/cli/main.go

.PHONY: build
