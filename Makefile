# Simple Makefile for a Go project

# Build the application
all: build_all

build:
	@echo "Building..."
	
	@go build -o us-solver cmd/main.go

# Run the application
run:
	@go run cmd/main.go

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f us-solver us-solver_*

build_windows_AMD64:
	@echo "Building for Windows (AMD64)..."
	@GOOS=windows GOARCH=amd64 go build -o us-solver_windows_amd64.exe cmd/main.go

build_windows_ARM64:
	@echo "Building for Windows (ARM64)..."
	@GOOS=windows GOARCH=arm64 go build -o us-solver_windows_arm64.exe cmd/main.go

build_linux_AMD64:
	@echo "Building for Linux (AMD64)..."
	@GOOS=linux GOARCH=amd64 go build -o us-solver_linux_amd64 cmd/main.go

build_linux_ARM64:
	@echo "Building for Linux (ARM64)..."
	@GOOS=linux GOARCH=arm64 go build -o us-solver_linux_arm64 cmd/main.go

build_mac_AMD64:
	@echo "Building for macOS (AMD64)..." 
	@GOOS=darwin GOARCH=amd64 go build -o us-solver_mac_amd64 cmd/main.go

build_mac_ARM64:
	@echo "Building for macOS (ARM64)..."
	GOOS=darwin GOARCH=arm64 go build -o us-solver_mac_silicon cmd/main.go
build_all: build_windows_AMD64 build_windows_ARM64 build_linux_AMD64 build_linux_ARM64 build_mac_AMD64 build_mac_ARM64

.PHONY: all build run test clean build_windows_AMD64 build_windows_ARM64 build_linux_AMD64 build_linux_ARM64 build_mac_AMD64 build_mac_ARM64
