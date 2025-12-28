package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Example 1: Worker Pool Pattern
type Job struct {
	ID   int
	Data string
}

type Result struct {
	Job    Job
	Output string
}

func workerPool(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for job := range jobs {
		// Simulate work
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		
		result := Result{
			Job:    job,
			Output: fmt.Sprintf("Worker %d processed job %d: %s", id, job.ID, job.Data),
		}
		results <- result
	}
}

func workerPoolExample() {
	fmt.Println("\n=== Example 1: Worker Pool Pattern ===")
	
	numWorkers := 3
	numJobs := 10
	
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	
	var wg sync.WaitGroup
	
	// Start workers
	fmt.Printf("Starting %d workers...\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go workerPool(w, jobs, results, &wg)
	}
	
	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Data: fmt.Sprintf("task-%d", j)}
	}
	close(jobs)
	
	// Close results when all workers done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	for result := range results {
		fmt.Println(result.Output)
	}
}

// Example 2: Fan-Out / Fan-In Pattern
func producer(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		ch <- i
	}
	close(ch)
}

func fanOutWorker(id int, input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for num := range input {
		// Process (square the number)
		result := num * num
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Worker %d: %d² = %d\n", id, num, result)
		output <- result
	}
}

func fanInFanOutExample() {
	fmt.Println("\n=== Example 2: Fan-Out / Fan-In Pattern ===")
	
	input := make(chan int)
	output := make(chan int)
	
	var wg sync.WaitGroup
	
	// Start producer
	go producer(input, 10)
	
	// Fan-out: Start multiple workers reading from same input channel
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go fanOutWorker(w, input, output, &wg)
	}
	
	// Fan-in: Close output channel when all workers done
	go func() {
		wg.Wait()
		close(output)
	}()
	
	// Collect all results
	fmt.Println("\nResults:")
	sum := 0
	for result := range output {
		sum += result
	}
	fmt.Printf("Sum of all results: %d\n", sum)
}

// Example 3: Context for Cancellation
func cancellableWorker(ctx context.Context, id int, results chan<- string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Received cancellation signal\n", id)
			results <- fmt.Sprintf("Worker %d cancelled", id)
			return
		default:
			// Do work
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Worker %d: Working...\n", id)
		}
	}
}

func contextCancellationExample() {
	fmt.Println("\n=== Example 3: Context Cancellation ===")
	
	// Create context with cancel function
	ctx, cancel := context.WithCancel(context.Background())
	results := make(chan string, 3)
	
	// Start workers
	for i := 1; i <= 3; i++ {
		go cancellableWorker(ctx, i, results)
	}
	
	// Let them work for a bit
	time.Sleep(500 * time.Millisecond)
	
	// Cancel all workers
	fmt.Println("\n🛑 Sending cancellation signal...")
	cancel()
	
	// Collect cancellation confirmations
	for i := 0; i < 3; i++ {
		msg := <-results
		fmt.Println(msg)
	}
}

// Example 4: Context with Timeout
func taskWithTimeout(ctx context.Context, id int, duration time.Duration) error {
	select {
	case <-time.After(duration):
		fmt.Printf("Task %d: Completed in %v\n", id, duration)
		return nil
	case <-ctx.Done():
		fmt.Printf("Task %d: Timeout after %v\n", id, duration)
		return ctx.Err()
	}
}

func contextTimeoutExample() {
	fmt.Println("\n=== Example 4: Context with Timeout ===")
	
	// Create context with 300ms timeout
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	
	// Launch tasks with different durations
	tasks := []time.Duration{100 * time.Millisecond, 200 * time.Millisecond, 400 * time.Millisecond}
	
	var wg sync.WaitGroup
	for i, duration := range tasks {
		wg.Add(1)
		go func(id int, d time.Duration) {
			defer wg.Done()
			taskWithTimeout(ctx, id, d)
		}(i+1, duration)
	}
	
	wg.Wait()
	fmt.Println("All tasks completed or timed out")
}

// Example 5: Rate Limiting with Ticker
func rateLimitedWorker(id int, requests <-chan int, ticker *time.Ticker, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for req := range requests {
		<-ticker.C // Wait for ticker (rate limit)
		fmt.Printf("Worker %d: Processing request %d at %v\n", id, req, time.Now().Format("15:04:05.000"))
	}
}

func rateLimitingExample() {
	fmt.Println("\n=== Example 5: Rate Limiting ===")
	
	requests := make(chan int, 10)
	ticker := time.NewTicker(200 * time.Millisecond) // Rate limit: 5 requests per second
	defer ticker.Stop()
	
	var wg sync.WaitGroup
	
	// Start worker
	wg.Add(1)
	go rateLimitedWorker(1, requests, ticker, &wg)
	
	// Send 5 requests
	fmt.Println("Sending 5 requests (rate limited to 1 per 200ms)...")
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	
	wg.Wait()
}

// Example 6: Semaphore Pattern (Limiting Concurrent Goroutines)
func semaphoreTask(id int, sem chan struct{}, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Acquire semaphore
	sem <- struct{}{}
	defer func() { <-sem }() // Release semaphore
	
	fmt.Printf("Task %d: Started (limited concurrency)\n", id)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	results <- fmt.Sprintf("Task %d completed", id)
}

func semaphoreExample() {
	fmt.Println("\n=== Example 6: Semaphore Pattern ===")
	fmt.Println("Limiting to max 2 concurrent tasks...")
	
	maxConcurrent := 2
	sem := make(chan struct{}, maxConcurrent) // Buffered channel as semaphore
	results := make(chan string, 5)
	
	var wg sync.WaitGroup
	
	// Launch 5 tasks, but only 2 will run concurrently
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go semaphoreTask(i, sem, results, &wg)
	}
	
	// Close results when all done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	for result := range results {
		fmt.Println(result)
	}
}

// Example 7: Error Group Pattern
type ErrorGroup struct {
	wg     sync.WaitGroup
	mu     sync.Mutex
	errors []error
}

func (eg *ErrorGroup) Go(f func() error) {
	eg.wg.Add(1)
	go func() {
		defer eg.wg.Done()
		if err := f(); err != nil {
			eg.mu.Lock()
			eg.errors = append(eg.errors, err)
			eg.mu.Unlock()
		}
	}()
}

func (eg *ErrorGroup) Wait() []error {
	eg.wg.Wait()
	return eg.errors
}

func errorGroupExample() {
	fmt.Println("\n=== Example 7: Error Group Pattern ===")
	
	var eg ErrorGroup
	
	// Launch tasks that might fail
	eg.Go(func() error {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Task 1: Success")
		return nil
	})
	
	eg.Go(func() error {
		time.Sleep(150 * time.Millisecond)
		fmt.Println("Task 2: Failed")
		return fmt.Errorf("task 2 error: something went wrong")
	})
	
	eg.Go(func() error {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Task 3: Success")
		return nil
	})
	
	// Wait and collect errors
	errors := eg.Wait()
	
	if len(errors) > 0 {
		fmt.Printf("\n❌ %d task(s) failed:\n", len(errors))
		for i, err := range errors {
			fmt.Printf("  %d. %v\n", i+1, err)
		}
	} else {
		fmt.Println("\n✅ All tasks completed successfully!")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║   Advanced Goroutine Patterns             ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	
	workerPoolExample()
	fanInFanOutExample()
	contextCancellationExample()
	contextTimeoutExample()
	rateLimitingExample()
	semaphoreExample()
	errorGroupExample()
	
	fmt.Println("\n✅ All advanced examples completed!")
	fmt.Println("\nAdvanced Patterns Summary:")
	fmt.Println("1. Worker Pool: Fixed number of workers processing jobs")
	fmt.Println("2. Fan-Out/Fan-In: Distribute work, collect results")
	fmt.Println("3. Context Cancellation: Graceful shutdown of goroutines")
	fmt.Println("4. Context Timeout: Automatic cancellation after time limit")
	fmt.Println("5. Rate Limiting: Control request rate with ticker")
	fmt.Println("6. Semaphore: Limit concurrent goroutines")
	fmt.Println("7. Error Group: Collect errors from multiple goroutines")
}
