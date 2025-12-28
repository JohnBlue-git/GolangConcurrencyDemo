package main

import (
	"fmt"
	"time"
)

// Example 1: The most basic goroutine
func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func basicExample() {
	fmt.Println("\n=== Example 1: Basic Goroutine ===")
	
	// Normal function call - runs synchronously
	fmt.Println("Before goroutine")
	
	// Launch a goroutine with 'go' keyword
	go sayHello()
	
	// Main function continues immediately
	fmt.Println("After launching goroutine")
	
	// Sleep to give goroutine time to execute
	// Without this, main() might exit before goroutine runs!
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("Main function ending")
}

// Example 2: Multiple goroutines
func printNumbers(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func multipleGoroutines() {
	fmt.Println("\n=== Example 2: Multiple Goroutines ===")
	
	// Launch multiple goroutines
	go printNumbers("Goroutine-1")
	go printNumbers("Goroutine-2")
	go printNumbers("Goroutine-3")
	
	// Notice how they run concurrently (interleaved output)
	time.Sleep(600 * time.Millisecond)
}

// Example 3: Goroutine with anonymous function
func anonymousGoroutine() {
	fmt.Println("\n=== Example 3: Anonymous Function Goroutine ===")
	
	message := "Hello from anonymous goroutine"
	
	// Launch goroutine with anonymous function
	go func() {
		fmt.Println(message)
		fmt.Println("This is an inline goroutine!")
	}()
	
	time.Sleep(100 * time.Millisecond)
}

// Example 4: Passing arguments to goroutines
func greet(name string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Printf("Hello, %s!\n", name)
}

func goroutineWithArgs() {
	fmt.Println("\n=== Example 4: Goroutines with Arguments ===")
	
	// Pass arguments to goroutine functions
	go greet("Alice", 100*time.Millisecond)
	go greet("Bob", 50*time.Millisecond)
	go greet("Charlie", 150*time.Millisecond)
	
	time.Sleep(200 * time.Millisecond)
	
	// Notice: Bob responds first (shortest delay)
}

// Example 5: Common pitfall - loop variable capture
func loopVariablePitfall() {
	fmt.Println("\n=== Example 5: Loop Variable Pitfall (WRONG) ===")
	
	// WRONG WAY - all goroutines will likely print the same value
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("Wrong: %d\n", i) // Captures loop variable by reference
		}()
	}
	
	time.Sleep(100 * time.Millisecond)
}

func loopVariableFixed() {
	fmt.Println("\n=== Example 5b: Loop Variable Fixed (CORRECT) ===")
	
	// CORRECT WAY 1 - Pass as argument
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Printf("Correct (arg): %d\n", n)
		}(i) // Pass i as argument
	}
	
	time.Sleep(100 * time.Millisecond)
	
	// CORRECT WAY 2 - Create local copy
	for i := 0; i < 5; i++ {
		i := i // Create a new variable for each iteration
		go func() {
			fmt.Printf("Correct (copy): %d\n", i)
		}()
	}
	
	time.Sleep(100 * time.Millisecond)
}

func main() {
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║   Basic Goroutine Examples                ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	
	basicExample()
	multipleGoroutines()
	anonymousGoroutine()
	goroutineWithArgs()
	loopVariablePitfall()
	loopVariableFixed()
	
	fmt.Println("\n✅ All basic examples completed!")
	fmt.Println("\nKey Takeaways:")
	fmt.Println("1. Use 'go' keyword to launch a goroutine")
	fmt.Println("2. Goroutines run concurrently, not sequentially")
	fmt.Println("3. Main function won't wait for goroutines automatically")
	fmt.Println("4. Be careful with loop variables - pass as arguments!")
	fmt.Println("5. Goroutines are very lightweight (~2KB each)")
}
