package main

import (
	"fmt"
	"sync"
	"time"
)

// Example 1: Using WaitGroup to wait for goroutines
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when goroutine completes
	
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id*100) * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

func waitGroupExample() {
	fmt.Println("\n=== Example 1: WaitGroup ===")
	
	var wg sync.WaitGroup
	
	// Launch 5 workers
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment counter before launching goroutine
		go worker(i, &wg)
	}
	
	fmt.Println("Waiting for all workers to complete...")
	wg.Wait() // Block until counter becomes 0
	fmt.Println("All workers completed!")
}

// Example 2: Basic channel communication
func sendData(ch chan string) {
	time.Sleep(100 * time.Millisecond)
	ch <- "Hello from goroutine!" // Send data to channel
}

func basicChannel() {
	fmt.Println("\n=== Example 2: Basic Channel ===")
	
	// Create a channel
	ch := make(chan string)
	
	// Launch goroutine that sends data
	go sendData(ch)
	
	// Receive data from channel (this blocks until data arrives)
	message := <-ch
	fmt.Println("Received:", message)
}

// Example 3: Buffered channels
func bufferedChannel() {
	fmt.Println("\n=== Example 3: Buffered Channel ===")
	
	// Create buffered channel (capacity 3)
	ch := make(chan int, 3)
	
	// Can send 3 values without blocking
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Sent 3 values to buffered channel")
	
	// Receive values
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
}

// Example 4: Channel direction (send-only, receive-only)
func sender(ch chan<- int) { // Send-only channel
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
	}
	close(ch) // Sender closes the channel when done
}

func receiver(ch <-chan int, done chan<- bool) { // Receive-only channel
	for num := range ch { // Loop until channel is closed
		fmt.Printf("Received: %d\n", num)
	}
	done <- true
}

func channelDirection() {
	fmt.Println("\n=== Example 4: Channel Direction ===")
	
	ch := make(chan int)
	done := make(chan bool)
	
	go sender(ch)
	go receiver(ch, done)
	
	<-done // Wait for receiver to finish
	fmt.Println("Communication completed!")
}

// Example 5: Multiple goroutines with WaitGroup and channels
func processor(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2 // Send result
	}
}

func pipelineExample() {
	fmt.Println("\n=== Example 5: Pipeline with WaitGroup ===")
	
	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	var wg sync.WaitGroup
	
	// Start 3 worker goroutines
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go processor(w, jobs, results, &wg)
	}
	
	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close jobs channel - no more jobs coming
	
	// Wait for all workers to finish, then close results
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	fmt.Println("\nResults:")
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// Example 6: Select statement for multiple channels
func selectExample() {
	fmt.Println("\n=== Example 6: Select Statement ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Two goroutines sending to different channels
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()
	
	// Select receives from whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout!")
		}
	}
}

// Example 7: Non-blocking channel operations
func nonBlockingChannel() {
	fmt.Println("\n=== Example 7: Non-blocking Channel Operations ===")
	
	messages := make(chan string)
	signals := make(chan bool)
	
	// Non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received (non-blocking)")
	}
	
	// Non-blocking send
	msg := "Hi there"
	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("No message sent (channel not ready)")
	}
	
	// Multiple non-blocking operations
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("No activity")
	}
}

func main() {
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║   Intermediate Goroutine Examples         ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	
	waitGroupExample()
	basicChannel()
	bufferedChannel()
	channelDirection()
	pipelineExample()
	selectExample()
	nonBlockingChannel()
	
	fmt.Println("\n✅ All intermediate examples completed!")
	fmt.Println("\nKey Concepts:")
	fmt.Println("1. WaitGroup: Wait for multiple goroutines to complete")
	fmt.Println("2. Channels: Type-safe way to communicate between goroutines")
	fmt.Println("3. Buffered Channels: Allow sending without immediate receiver")
	fmt.Println("4. Channel Direction: Enforce send-only or receive-only")
	fmt.Println("5. close(): Sender closes channel when done sending")
	fmt.Println("6. range: Loop over channel until it's closed")
	fmt.Println("7. select: Handle multiple channel operations")
}
