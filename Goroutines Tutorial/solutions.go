package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
╔════════════════════════════════════════════════════════╗
║         GOROUTINES EXERCISE SOLUTIONS                  ║
╚════════════════════════════════════════════════════════╝

Solutions to exercises.go
Study these after attempting the exercises yourself!
*/

// ============================================
// SOLUTION 1: First Goroutine
// ============================================

func solution1() {
	fmt.Println("\n=== Solution 1: First Goroutine ===")
	
	// Create a simple function
	sayHello := func() {
		fmt.Println("Hello from goroutine!")
	}
	
	// Launch as goroutine
	go sayHello()
	
	// Wait for goroutine to execute
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("✅ Goroutine completed!")
}

// ============================================
// SOLUTION 2: Multiple Goroutines
// ============================================

func solution2() {
	fmt.Println("\n=== Solution 2: Multiple Goroutines ===")
	
	// Function that prints ID
	printID := func(id int) {
		fmt.Printf("Goroutine %d is running\n", id)
	}
	
	// Launch 5 goroutines
	for i := 1; i <= 5; i++ {
		go printID(i)
	}
	
	// Wait for all to complete
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("✅ All goroutines completed!")
}

// ============================================
// SOLUTION 3: Using WaitGroup
// ============================================

func solution3() {
	fmt.Println("\n=== Solution 3: Using WaitGroup ===")
	
	var wg sync.WaitGroup
	
	// Worker function
	worker := func(id int) {
		defer wg.Done() // Always defer Done()
		
		fmt.Printf("Worker %d: Starting\n", id)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("Worker %d: Done\n", id)
	}
	
	// Launch 10 goroutines
	for i := 1; i <= 10; i++ {
		wg.Add(1) // Add before launching
		go worker(i)
	}
	
	// Wait for all to complete
	wg.Wait()
	
	fmt.Println("✅ All workers completed!")
}

// ============================================
// SOLUTION 4: Channel Communication
// ============================================

func solution4() {
	fmt.Println("\n=== Solution 4: Channel Communication ===")
	
	// Create channel
	ch := make(chan int)
	
	// Sender goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i // Send to channel
		}
		close(ch) // Close when done
	}()
	
	// Receiver goroutine
	go func() {
		for num := range ch { // Receive until closed
			fmt.Printf("Received: %d\n", num)
		}
	}()
	
	// Wait for completion
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("✅ Channel communication completed!")
}

// ============================================
// SOLUTION 5: Parallel Sum
// ============================================

func solution5() {
	fmt.Println("\n=== Solution 5: Parallel Sum ===")
	
	// Create channel for partial sums
	ch := make(chan int)
	
	// Calculate sum function
	calculateSum := func(start, end int) {
		sum := 0
		for i := start; i <= end; i++ {
			sum += i
		}
		ch <- sum // Send partial sum
	}
	
	// Launch two goroutines
	go calculateSum(1, 50)   // Sum 1-50
	go calculateSum(51, 100) // Sum 51-100
	
	// Receive partial sums
	sum1 := <-ch
	sum2 := <-ch
	
	// Calculate total
	total := sum1 + sum2
	
	fmt.Printf("Sum 1-50: %d\n", sum1)
	fmt.Printf("Sum 51-100: %d\n", sum2)
	fmt.Printf("Total sum 1-100: %d\n", total)
	
	fmt.Println("✅ Parallel sum completed!")
}

// ============================================
// SOLUTION 6: Worker Pool
// ============================================

func solution6() {
	fmt.Println("\n=== Solution 6: Worker Pool ===")
	
	numWorkers := 3
	numJobs := 10
	
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	var wg sync.WaitGroup
	
	// Worker function
	worker := func(id int) {
		defer wg.Done()
		
		for job := range jobs {
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(100 * time.Millisecond)
			results <- job * 2 // Simple processing
		}
	}
	
	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w)
	}
	
	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Close results when all workers done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("\nResults:")
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
	
	fmt.Println("✅ Worker pool completed!")
}

// ============================================
// SOLUTION 7: Select Statement
// ============================================

func solution7() {
	fmt.Println("\n=== Solution 7: Select Statement ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// First sender (fast)
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()
	
	// Second sender (slow)
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()
	
	// Receive from both channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		case <-time.After(2 * time.Second):
			fmt.Println("⏱️ Timeout!")
		}
	}
	
	fmt.Println("✅ Select statement completed!")
}

// ============================================
// SOLUTION 8: Fix Race Condition
// ============================================

func solution8() {
	fmt.Println("\n=== Solution 8: Fix Race Condition ===")
	
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex // Add mutex
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()     // Lock before accessing shared variable
				counter++
				mu.Unlock()   // Unlock after
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("Counter: %d (expected: 10000)\n", counter)
	
	if counter == 10000 {
		fmt.Println("✅ Race condition fixed!")
	} else {
		fmt.Println("❌ Still has race condition!")
	}
}

// ============================================
// SOLUTION 9: Context Cancellation
// ============================================

func solution9() {
	fmt.Println("\n=== Solution 9: Context Cancellation ===")
	
	// Create context with cancel
	ctx, cancel := context.WithCancel(context.Background())
	
	// Long-running worker
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker: Received cancellation signal, stopping...")
				return
			default:
				fmt.Println("Worker: Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	
	// Let it work for 1 second
	time.Sleep(1 * time.Second)
	
	// Cancel the context
	fmt.Println("Main: Sending cancellation signal...")
	cancel()
	
	// Give time to see cancellation message
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("✅ Context cancellation completed!")
}

// ============================================
// SOLUTION 10: Pipeline
// ============================================

func solution10() {
	fmt.Println("\n=== Solution 10: Pipeline ===")
	
	// Stage 1: Generator
	generator := func(nums chan<- int) {
		for i := 1; i <= 10; i++ {
			nums <- i
		}
		close(nums)
	}
	
	// Stage 2: Squarer
	squarer := func(nums <-chan int, squares chan<- int) {
		for num := range nums {
			squares <- num * num
		}
		close(squares)
	}
	
	// Stage 3: Printer
	printer := func(squares <-chan int) {
		for square := range squares {
			fmt.Printf("%d ", square)
		}
		fmt.Println()
	}
	
	// Create channels
	nums := make(chan int)
	squares := make(chan int)
	
	// Connect stages
	go generator(nums)
	go squarer(nums, squares)
	printer(squares) // Run in main goroutine
	
	fmt.Println("✅ Pipeline completed!")
}

// ============================================
// BONUS SOLUTION: Parallel Web Fetcher
// ============================================

func bonusSolution() {
	fmt.Println("\n=== Bonus: Parallel Web Fetcher ===")
	
	urls := []string{
		"http://example.com/page1",
		"http://example.com/page2",
		"http://example.com/page3",
		"http://example.com/page4",
		"http://example.com/page5",
	}
	
	var wg sync.WaitGroup
	startTime := time.Now()
	
	// Fetch function
	fetch := func(url string) {
		defer wg.Done()
		
		// Simulate random fetch time (100-500ms)
		fetchTime := time.Duration(100+rand.Intn(400)) * time.Millisecond
		time.Sleep(fetchTime)
		
		fmt.Printf("Fetched %s in %v\n", url, fetchTime)
	}
	
	// Launch concurrent fetches
	for _, url := range urls {
		wg.Add(1)
		go fetch(url)
	}
	
	// Wait for all fetches
	wg.Wait()
	
	totalTime := time.Since(startTime)
	fmt.Printf("\n⏱️ Total time: %v (concurrent, not sequential!)\n", totalTime)
	fmt.Println("✅ Parallel fetching completed!")
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("╔════════════════════════════════════════════════════════╗")
	fmt.Println("║         GOROUTINES EXERCISE SOLUTIONS                  ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	fmt.Println("\nThese are the solutions to exercises.go")
	fmt.Println("Study them after attempting the exercises yourself!\n")
	
	solution1()
	time.Sleep(200 * time.Millisecond)
	
	solution2()
	time.Sleep(200 * time.Millisecond)
	
	solution3()
	time.Sleep(200 * time.Millisecond)
	
	solution4()
	time.Sleep(200 * time.Millisecond)
	
	solution5()
	time.Sleep(200 * time.Millisecond)
	
	solution6()
	time.Sleep(200 * time.Millisecond)
	
	solution7()
	time.Sleep(200 * time.Millisecond)
	
	solution8()
	time.Sleep(200 * time.Millisecond)
	
	solution9()
	time.Sleep(200 * time.Millisecond)
	
	solution10()
	time.Sleep(200 * time.Millisecond)
	
	bonusSolution()
	
	fmt.Println("\n╔════════════════════════════════════════════════════════╗")
	fmt.Println("║  All solutions demonstrated! Keep practicing! 🚀       ║")
	fmt.Println("╚════════════════════════════════════════════════════════╝")
	
	fmt.Println("\n📚 Key Takeaways:")
	fmt.Println("1. Use WaitGroup for synchronization")
	fmt.Println("2. Channels enable safe communication")
	fmt.Println("3. Always close channels when done sending")
	fmt.Println("4. Use mutex for shared variable access")
	fmt.Println("5. Context enables graceful cancellation")
	fmt.Println("6. Select handles multiple channel operations")
	fmt.Println("7. Pipelines connect processing stages")
	fmt.Println("8. Worker pools control resource usage")
}
