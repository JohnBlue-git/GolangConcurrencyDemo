package main

import (
	"fmt"
	"time"
)

/*
╔════════════════════════════════════════════════════════╗
║           GOROUTINES PRACTICE EXERCISES                ║
╚════════════════════════════════════════════════════════╝

Complete each exercise below. Solutions are in solutions.go
Run this file: go run exercises.go
*/

// ============================================
// EXERCISE 1: First Goroutine (Easy)
// ============================================
// TODO: Create a function that prints "Hello from goroutine!"
//       Launch it as a goroutine from main
//       Make sure it has time to execute before main exits

func exercise1() {
	fmt.Println("\n=== Exercise 1: First Goroutine ===")
	
	// Your code here:
	// 1. Create a function that prints a message
	// 2. Launch it with 'go'
	// 3. Wait for it to complete
	
	fmt.Println("Exercise 1 not implemented yet!")
}

// ============================================
// EXERCISE 2: Multiple Goroutines (Easy)
// ============================================
// TODO: Launch 5 goroutines that each print their ID
//       They should all run concurrently
//       Wait for all to complete

func exercise2() {
	fmt.Println("\n=== Exercise 2: Multiple Goroutines ===")
	
	// Your code here:
	// 1. Create a function that takes an ID parameter
	// 2. Launch 5 goroutines with different IDs
	// 3. Ensure all complete before function returns
	
	fmt.Println("Exercise 2 not implemented yet!")
}

// ============================================
// EXERCISE 3: Using WaitGroup (Medium)
// ============================================
// TODO: Create 10 goroutines using WaitGroup for synchronization
//       Each goroutine should simulate work with time.Sleep
//       Print start and end messages with goroutine ID

func exercise3() {
	fmt.Println("\n=== Exercise 3: Using WaitGroup ===")
	
	// Your code here:
	// 1. Create a sync.WaitGroup
	// 2. Launch 10 goroutines
	// 3. Each should call wg.Done() when finished
	// 4. Wait for all to complete with wg.Wait()
	
	fmt.Println("Exercise 3 not implemented yet!")
}

// ============================================
// EXERCISE 4: Channel Communication (Medium)
// ============================================
// TODO: Create two goroutines that communicate via a channel
//       First goroutine: sends numbers 1-5 to channel
//       Second goroutine: receives and prints numbers

func exercise4() {
	fmt.Println("\n=== Exercise 4: Channel Communication ===")
	
	// Your code here:
	// 1. Create a channel: ch := make(chan int)
	// 2. First goroutine sends numbers 1-5, then closes channel
	// 3. Second goroutine receives with 'range' and prints
	
	fmt.Println("Exercise 4 not implemented yet!")
}

// ============================================
// EXERCISE 5: Sum with Channels (Medium)
// ============================================
// TODO: Create a function that calculates sum of numbers 1-100
//       Split work between 2 goroutines:
//       - First: sum 1-50
//       - Second: sum 51-100
//       Use channels to send partial sums, then add them

func exercise5() {
	fmt.Println("\n=== Exercise 5: Parallel Sum ===")
	
	// Your code here:
	// 1. Create a channel for integers
	// 2. Launch 2 goroutines, each calculates partial sum
	// 3. Receive both partial sums and add them
	// 4. Print final result (should be 5050)
	
	fmt.Println("Exercise 5 not implemented yet!")
}

// ============================================
// EXERCISE 6: Worker Pool (Hard)
// ============================================
// TODO: Implement a worker pool with 3 workers
//       Create 10 jobs (just print job number)
//       Workers should process jobs concurrently

func exercise6() {
	fmt.Println("\n=== Exercise 6: Worker Pool ===")
	
	// Your code here:
	// 1. Create jobs channel: jobs := make(chan int, 10)
	// 2. Create results channel
	// 3. Start 3 worker goroutines
	// 4. Send 10 jobs to jobs channel
	// 5. Collect and print results
	
	fmt.Println("Exercise 6 not implemented yet!")
}

// ============================================
// EXERCISE 7: Select Statement (Hard)
// ============================================
// TODO: Create two channels with different message timings
//       Use select to receive from whichever is ready first
//       Implement a 2-second timeout

func exercise7() {
	fmt.Println("\n=== Exercise 7: Select Statement ===")
	
	// Your code here:
	// 1. Create two channels
	// 2. Launch goroutines that send to channels with delays
	// 3. Use select to receive from either channel
	// 4. Add timeout case: case <-time.After(2*time.Second)
	
	fmt.Println("Exercise 7 not implemented yet!")
}

// ============================================
// EXERCISE 8: Race Condition Fix (Hard)
// ============================================
// TODO: The code below has a race condition
//       Fix it using sync.Mutex
//       Final counter should be exactly 10000

func exercise8() {
	fmt.Println("\n=== Exercise 8: Fix Race Condition ===")
	
	// Broken code (race condition):
	/*
	counter := 0
	var wg sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++ // Race condition!
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("Counter: %d (should be 10000)\n", counter)
	*/
	
	// Your code here:
	// 1. Add sync.Mutex
	// 2. Lock before incrementing counter
	// 3. Unlock after incrementing
	// 4. Run with: go run -race exercises.go
	
	fmt.Println("Exercise 8 not implemented yet!")
}

// ============================================
// EXERCISE 9: Context Cancellation (Advanced)
// ============================================
// TODO: Create a long-running goroutine
//       Use context to cancel it after 1 second
//       Goroutine should detect cancellation and stop

func exercise9() {
	fmt.Println("\n=== Exercise 9: Context Cancellation ===")
	
	// Your code here:
	// 1. Create context: ctx, cancel := context.WithCancel(...)
	// 2. Launch goroutine that checks ctx.Done()
	// 3. After 1 second, call cancel()
	// 4. Goroutine should exit gracefully
	
	fmt.Println("Exercise 9 not implemented yet!")
}

// ============================================
// EXERCISE 10: Pipeline (Advanced)
// ============================================
// TODO: Build a 3-stage pipeline:
//       Stage 1: Generate numbers 1-10
//       Stage 2: Square each number
//       Stage 3: Print results
//       Connect stages with channels

func exercise10() {
	fmt.Println("\n=== Exercise 10: Pipeline ===")
	
	// Your code here:
	// 1. Create channels to connect stages
	// 2. Stage 1: send numbers 1-10, close channel
	// 3. Stage 2: receive, square, send result, close output
	// 4. Stage 3: receive and print
	
	fmt.Println("Exercise 10 not implemented yet!")
}

// ============================================
// BONUS EXERCISE: Parallel Web Fetcher
// ============================================
// TODO: Simulate fetching 5 URLs concurrently
//       Each "fetch" should sleep for random duration (100-500ms)
//       Print total time taken (should be ~500ms, not 1500ms)

func bonusExercise() {
	fmt.Println("\n=== Bonus: Parallel Web Fetcher ===")
	
	// Your code here:
	// 1. Create slice of "URLs" (just strings)
	// 2. Launch goroutine for each URL
	// 3. Simulate fetch with random sleep
	// 4. Use WaitGroup to wait for all
	// 5. Measure total time with time.Now() and time.Since()
	
	fmt.Println("Bonus exercise not implemented yet!")
}

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         GOROUTINES PRACTICE EXERCISES                  ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	fmt.Println("\nInstructions:")
	fmt.Println("1. Complete each exercise function above")
	fmt.Println("2. Run: go run exercises.go")
	fmt.Println("3. Check solutions: solutions.go")
	fmt.Println("4. Test for race conditions: go run -race exercises.go")
	
	// Uncomment exercises as you complete them:
	
	exercise1()
	time.Sleep(200 * time.Millisecond)
	
	exercise2()
	time.Sleep(200 * time.Millisecond)
	
	exercise3()
	time.Sleep(200 * time.Millisecond)
	
	exercise4()
	time.Sleep(200 * time.Millisecond)
	
	exercise5()
	time.Sleep(200 * time.Millisecond)
	
	exercise6()
	time.Sleep(200 * time.Millisecond)
	
	exercise7()
	time.Sleep(200 * time.Millisecond)
	
	exercise8()
	time.Sleep(200 * time.Millisecond)
	
	exercise9()
	time.Sleep(200 * time.Millisecond)
	
	exercise10()
	time.Sleep(200 * time.Millisecond)
	
	bonusExercise()
	
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║  Great job practicing! Check solutions.go for answers  ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
}
