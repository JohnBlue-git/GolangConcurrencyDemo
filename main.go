package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker represents a worker that processes jobs
type Worker struct {
	id int
}

// Process simulates work being done by a worker
func (w Worker) Process(job int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	
	// Simulate some work with random duration
	processingTime := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(processingTime)
	
	result := fmt.Sprintf("Worker %d processed job %d in %v", w.id, job, processingTime)
	results <- result
}

// fetchData simulates an asynchronous data fetch operation
func fetchData(source string, wg *sync.WaitGroup, dataChan chan<- string) {
	defer wg.Done()
	
	// Simulate network delay
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	
	data := fmt.Sprintf("Data fetched from %s", source)
	dataChan <- data
}

// counter demonstrates safe concurrent counter using mutex
func counter(name string, iterations int, mu *sync.Mutex, sharedCounter *int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for i := 0; i < iterations; i++ {
		mu.Lock()
		*sharedCounter++
		current := *sharedCounter
		mu.Unlock()
		
		if i%100 == 0 {
			fmt.Printf("%s: Counter at %d\n", name, current)
		}
		time.Sleep(time.Millisecond)
	}
}

// pipelineStage1 - first stage of pipeline
func pipelineStage1(input <-chan int, output chan<- int) {
	for num := range input {
		// Square the number
		result := num * num
		output <- result
	}
	close(output)
}

// pipelineStage2 - second stage of pipeline
func pipelineStage2(input <-chan int, output chan<- string) {
	for num := range input {
		// Convert to string with formatting
		result := fmt.Sprintf("Result: %d", num)
		output <- result
	}
	close(output)
}

// demonstrateWorkerPool shows concurrent worker pool pattern
func demonstrateWorkerPool() {
	fmt.Println("\n=== Worker Pool Demo ===")
	
	numWorkers := 5
	numJobs := 15
	
	var wg sync.WaitGroup
	results := make(chan string, numJobs)
	
	// Launch workers
	for i := 1; i <= numJobs; i++ {
		wg.Add(1)
		worker := Worker{id: (i % numWorkers) + 1}
		go worker.Process(i, &wg, results)
	}
	
	// Wait for all workers to complete and close results channel
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect and print results
	for result := range results {
		fmt.Println(result)
	}
}

// demonstrateAsyncFetching shows async data fetching pattern
func demonstrateAsyncFetching() {
	fmt.Println("\n=== Async Data Fetching Demo ===")
	
	sources := []string{"API-1", "API-2", "API-3", "Database", "Cache"}
	
	var wg sync.WaitGroup
	dataChan := make(chan string, len(sources))
	
	// Launch async fetch operations
	startTime := time.Now()
	for _, source := range sources {
		wg.Add(1)
		go fetchData(source, &wg, dataChan)
	}
	
	// Wait for all fetches to complete
	go func() {
		wg.Wait()
		close(dataChan)
	}()
	
	// Collect results
	for data := range dataChan {
		fmt.Println(data)
	}
	
	fmt.Printf("Total time: %v\n", time.Since(startTime))
}

// demonstrateMutex shows thread-safe counter using mutex
func demonstrateMutex() {
	fmt.Println("\n=== Mutex Demo (Thread-Safe Counter) ===")
	
	var mu sync.Mutex
	var sharedCounter int
	var wg sync.WaitGroup
	
	// Launch multiple goroutines that increment a shared counter
	goroutines := []string{"Goroutine-A", "Goroutine-B", "Goroutine-C"}
	
	for _, name := range goroutines {
		wg.Add(1)
		go counter(name, 200, &mu, &sharedCounter, &wg)
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", sharedCounter)
}

// demonstratePipeline shows channel pipeline pattern
func demonstratePipeline() {
	fmt.Println("\n=== Pipeline Demo ===")
	
	// Create channels for pipeline stages
	stage1Input := make(chan int)
	stage1Output := make(chan int)
	stage2Output := make(chan string)
	
	// Start pipeline stages
	go pipelineStage1(stage1Input, stage1Output)
	go pipelineStage2(stage1Output, stage2Output)
	
	// Send data into pipeline
	go func() {
		for i := 1; i <= 5; i++ {
			stage1Input <- i
		}
		close(stage1Input)
	}()
	
	// Receive results from pipeline
	for result := range stage2Output {
		fmt.Println(result)
	}
}

// demonstrateSelect shows select statement for channel multiplexing
func demonstrateSelect() {
	fmt.Println("\n=== Select Statement Demo ===")
	
	chan1 := make(chan string)
	chan2 := make(chan string)
	
	// Launch two goroutines sending to different channels
	go func() {
		time.Sleep(300 * time.Millisecond)
		chan1 <- "Message from Channel 1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		chan2 <- "Message from Channel 2"
	}()
	
	// Use select to receive from whichever channel is ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("Received:", msg1)
		case msg2 := <-chan2:
			fmt.Println("Received:", msg2)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout!")
		}
	}
}

// ============================================================================
// INTEGRATED DEMO: Combining ALL Patterns Together
// ============================================================================

// Simulated API data
type APIResponse struct {
	Source string
	Data   string
	Time   time.Duration
}

// Processed result
type ProcessedData struct {
	ID        int
	Original  string
	Processed string
	Source    string
}

// Statistics (protected by mutex)
type Stats struct {
	mu           sync.Mutex
	totalFetched int
	totalProcessed int
	errors       int
}

func (s *Stats) IncrementFetched() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.totalFetched++
}

func (s *Stats) IncrementProcessed() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.totalProcessed++
}

func (s *Stats) IncrementErrors() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.errors++
}

func (s *Stats) Print() {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("\n📊 Final Statistics:\n")
	fmt.Printf("   Fetched: %d | Processed: %d | Errors: %d\n", 
		s.totalFetched, s.totalProcessed, s.errors)
}

// Stage 1: Async fetching with timeout (using select)
func fetchWithTimeout(source string, timeout time.Duration, stats *Stats) (*APIResponse, error) {
	responseChan := make(chan *APIResponse, 1)
	errorChan := make(chan error, 1)
	
	// Simulate async fetch
	go func() {
		fetchTime := time.Duration(rand.Intn(800)) * time.Millisecond
		time.Sleep(fetchTime)
		
		responseChan <- &APIResponse{
			Source: source,
			Data:   fmt.Sprintf("data-from-%s", source),
			Time:   fetchTime,
		}
	}()
	
	// Use SELECT to handle timeout
	select {
	case response := <-responseChan:
		stats.IncrementFetched()
		return response, nil
	case err := <-errorChan:
		stats.IncrementErrors()
		return nil, err
	case <-time.After(timeout):
		stats.IncrementErrors()
		return nil, fmt.Errorf("timeout fetching from %s", source)
	}
}

// Stage 2: Worker pool for processing
func processingWorker(id int, jobs <-chan *APIResponse, results chan<- ProcessedData, 
	stats *Stats, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for job := range jobs {
		// Simulate processing
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		
		result := ProcessedData{
			ID:        id,
			Original:  job.Data,
			Processed: fmt.Sprintf("PROCESSED[%s]", job.Data),
			Source:    job.Source,
		}
		
		stats.IncrementProcessed()
		results <- result
	}
}

// Stage 3: Pipeline for final output
func outputPipeline(results <-chan ProcessedData, done chan<- bool) {
	fmt.Println("\n📦 Processing Results:")
	count := 0
	for result := range results {
		fmt.Printf("   Worker-%d: %s (from %s)\n", 
			result.ID, result.Processed, result.Source)
		count++
	}
	fmt.Printf("   Total results: %d\n", count)
	done <- true
}

// INTEGRATED DEMONSTRATION
func demonstrateIntegrated() {
	fmt.Println("\n=== 🎯 INTEGRATED DEMO: All Patterns Combined ===")
	fmt.Println("Scenario: Fetch data from APIs, process with workers, output via pipeline")
	fmt.Println()
	
	// Initialize statistics (MUTEX pattern)
	stats := &Stats{}
	
	// Data sources
	sources := []string{"API-1", "API-2", "API-3", "API-4", "API-5"}
	
	// Channels for pipeline
	fetchedData := make(chan *APIResponse, len(sources))
	processedData := make(chan ProcessedData, len(sources))
	done := make(chan bool)
	
	// ========================================
	// STAGE 1: ASYNC FETCHING with SELECT
	// ========================================
	fmt.Println("🌐 Stage 1: Fetching from multiple APIs concurrently...")
	var fetchWg sync.WaitGroup
	
	for _, source := range sources {
		fetchWg.Add(1)
		go func(src string) {
			defer fetchWg.Done()
			
			// Fetch with 1 second timeout (SELECT pattern)
			response, err := fetchWithTimeout(src, 1*time.Second, stats)
			if err != nil {
				fmt.Printf("   ⚠️  Error: %v\n", err)
				return
			}
			
			fmt.Printf("   ✓ Fetched from %s in %v\n", response.Source, response.Time)
			fetchedData <- response
		}(source)
	}
	
	// Wait for all fetches, then close channel
	go func() {
		fetchWg.Wait()
		close(fetchedData)
		fmt.Println("   All fetches complete!\n")
	}()
	
	// ========================================
	// STAGE 2: WORKER POOL for processing
	// ========================================
	fmt.Println("⚙️  Stage 2: Processing data with worker pool...")
	var processWg sync.WaitGroup
	numWorkers := 3
	
	// Start workers
	for w := 1; w <= numWorkers; w++ {
		processWg.Add(1)
		go processingWorker(w, fetchedData, processedData, stats, &processWg)
	}
	
	// Close results when all workers done
	go func() {
		processWg.Wait()
		close(processedData)
	}()
	
	// ========================================
	// STAGE 3: PIPELINE for output
	// ========================================
	go outputPipeline(processedData, done)
	
	// Wait for pipeline to complete
	<-done
	
	// Print statistics (MUTEX protected)
	stats.Print()
	
	fmt.Println("\n✅ Integrated demo completed!")
	fmt.Println("\nPatterns used:")
	fmt.Println("   ✓ Async Fetching: Concurrent API calls")
	fmt.Println("   ✓ Select: Timeout handling")
	fmt.Println("   ✓ Worker Pool: Limited concurrent processors")
	fmt.Println("   ✓ Pipeline: Data flows through stages")
	fmt.Println("   ✓ Mutex: Thread-safe statistics")
	fmt.Println("   ✓ WaitGroups: Synchronization at each stage")
}

func main() {
	fmt.Println("===========================================")
	fmt.Println("  Go Concurrency & Async Programming Demo")
	fmt.Println("===========================================")
	
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	
	// Option 1: Individual demonstrations
	fmt.Println("\n📚 Choose demo mode:")
	fmt.Println("   Running INTEGRATED demo (combines all patterns)")
	fmt.Println()
	
	demonstrateIntegrated()
	
	fmt.Println("\n================================================")
	fmt.Println("💡 Want to see individual patterns?")
	fmt.Println("   Uncomment the sections below in main()")
	fmt.Println("================================================")
	
	// Uncomment to see individual demos:
	/*
	demonstrateWorkerPool()
	demonstrateAsyncFetching()
	demonstrateMutex()
	demonstratePipeline()
	demonstrateSelect()
	*/
	
	fmt.Println("\n===========================================")
	fmt.Println("  All demos completed successfully!")
	fmt.Println("===========================================")
}
