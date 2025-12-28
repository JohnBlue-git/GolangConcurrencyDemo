#!/bin/bash

# Quick Go Installation Script for Ubuntu/Debian
# This script will install Go and run the demo

echo "=== Installing Go ==="
sudo apt update
sudo apt install -y golang-go

echo ""
echo "=== Verifying Go Installation ==="
go version

echo ""
echo "=== Running the Concurrency Demo ==="
cd "$(dirname "$0")"
go run main.go

echo ""
echo "=== Installation and Demo Complete! ==="
echo "You can also build the project with: go build -o concurrency-demo"
