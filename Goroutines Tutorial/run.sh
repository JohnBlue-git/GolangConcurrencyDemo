#!/bin/bash

# Goroutines Tutorial Runner
# This script helps you run all the goroutine examples

echo "╔════════════════════════════════════════════════════════╗"
echo "║         GOROUTINES TUTORIAL - Quick Start             ║"
echo "╚════════════════════════════════════════════════════════╝"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed!"
    echo "📦 Install with: sudo apt install golang-go"
    echo "   Or: sudo snap install go --classic"
    exit 1
fi

echo "✅ Go version: $(go version)"
echo ""
echo "Available examples:"
echo "  1. Basic Goroutines (01_basic_goroutine.go)"
echo "  2. Intermediate Patterns (02_intermediate_goroutine.go)"
echo "  3. Advanced Patterns (03_advanced_goroutine.go)"
echo "  4. Practice Exercises (exercises.go)"
echo "  5. Exercise Solutions (solutions.go)"
echo "  A. Run all examples"
echo "  Q. Quit"
echo ""

read -p "Select an option (1-5, A, or Q): " choice

case $choice in
    1)
        echo ""
        echo "Running Basic Goroutines..."
        echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
        go run 01_basic_goroutine.go
        ;;
    2)
        echo ""
        echo "Running Intermediate Patterns..."
        echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
        go run 02_intermediate_goroutine.go
        ;;
    3)
        echo ""
        echo "Running Advanced Patterns..."
        echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
        go run 03_advanced_goroutine.go
        ;;
    4)
        echo ""
        echo "Running Practice Exercises..."
        echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
        echo "💡 Complete the exercises in exercises.go first!"
        go run exercises.go
        ;;
    5)
        echo ""
        echo "Running Exercise Solutions..."
        echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
        go run solutions.go
        ;;
    [Aa])
        echo ""
        echo "Running all examples..."
        echo ""
        
        echo "1️⃣  BASIC GOROUTINES"
        echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
        go run 01_basic_goroutine.go
        echo ""
        
        echo "2️⃣  INTERMEDIATE PATTERNS"
        echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
        go run 02_intermediate_goroutine.go
        echo ""
        
        echo "3️⃣  ADVANCED PATTERNS"
        echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
        go run 03_advanced_goroutine.go
        echo ""
        
        echo "4️⃣  SOLUTIONS"
        echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
        go run solutions.go
        ;;
    [Qq])
        echo "Goodbye! Happy learning! 🚀"
        exit 0
        ;;
    *)
        echo "❌ Invalid option. Please select 1-5, A, or Q"
        exit 1
        ;;
esac

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ Completed!"
echo ""
echo "💡 Tips:"
echo "  • Read the README for detailed explanations"
echo "  • Modify the examples to experiment"
echo "  • Run with race detector: go run -race <file>.go"
echo "  • Complete exercises.go for practice"
echo ""
echo "📚 Next steps:"
echo "  • Try the exercises in exercises.go"
echo "  • Read 'Effective Go' concurrency section"
echo "  • Build a small concurrent project"
