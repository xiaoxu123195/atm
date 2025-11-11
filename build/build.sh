#!/bin/bash

echo "========================================"
echo "Building ATM-Go"
echo "========================================"

# Create bin directory if it doesn't exist
mkdir -p ../bin

# Detect OS
OS=$(uname -s)
case "$OS" in
    Linux*)     BINARY_NAME=atm-linux;;
    Darwin*)    BINARY_NAME=atm-darwin;;
    *)          BINARY_NAME=atm;;
esac

# Build
echo "Building $BINARY_NAME..."
cd ..
go build -ldflags="-s -w" -o bin/$BINARY_NAME cmd/atm/main.go

if [ $? -eq 0 ]; then
    echo ""
    echo "========================================"
    echo "Build successful!"
    echo "========================================"
    echo ""
    echo "Executable created at: bin/$BINARY_NAME"
    echo ""
    echo "To install globally:"
    echo "  sudo cp bin/$BINARY_NAME /usr/local/bin/atm"
    echo "  sudo chmod +x /usr/local/bin/atm"
    echo ""
    echo "========================================"
else
    echo ""
    echo "========================================"
    echo "Build failed!"
    echo "========================================"
    exit 1
fi
