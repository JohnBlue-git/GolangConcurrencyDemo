# Go Concurrency & Async Programming Demo

## Introduction to Go (Golang)

Go, also known as Golang, is an open-source programming language developed by Google in 2007 and publicly released in 2009. Created by Robert Griesemer, Rob Pike, and Ken Thompson, Go was designed to address the challenges of modern software development.

### Why Go?

- **Simple and Clean Syntax**: Easy to learn and read, with a focus on simplicity
- **Fast Compilation**: Compiles to native machine code quickly
- **Built-in Concurrency**: First-class support for concurrent programming through goroutines and channels
- **Efficient Garbage Collection**: Automatic memory management with low-latency GC
- **Strong Standard Library**: Comprehensive built-in packages for common tasks
- **Static Typing**: Type safety with type inference for cleaner code
- **Cross-Platform**: Compile once, run anywhere (Windows, Linux, macOS, etc.)

### Go's Concurrency Model

Go's concurrency model is based on **CSP (Communicating Sequential Processes)**:

- **Goroutines**: Lightweight threads managed by the Go runtime (not OS threads)
  - Start with ~2KB stack size (vs ~1MB for OS threads)
  - Multiplexed onto multiple OS threads
  - Can easily run millions of goroutines concurrently

- **Channels**: Type-safe pipes for communication between goroutines
  - Allows goroutines to send and receive values
  - Provides synchronization without explicit locks

- **Philosophy**: *"Don't communicate by sharing memory; share memory by communicating"*

## Project Overview

This project demonstrates Go's powerful concurrency and asynchronous programming features through practical examples. It showcases various concurrency patterns commonly used in real-world Go applications.

### 🎯 Main Demo: Integrated Example

The main program (`main.go`) runs an **integrated demonstration** that combines ALL concurrency patterns together in a realistic scenario:

**Scenario**: Fetch data from multiple APIs, process with worker pool, output via pipeline

```
Stage 1: Async Fetching (with timeout)
   API-1, API-2, API-3, API-4, API-5 (all fetched concurrently)
          ↓
Stage 2: Worker Pool (3 workers processing data)
   Worker-1, Worker-2, Worker-3 (process incoming data)
          ↓
Stage 3: Pipeline (output results)
   Results → Console
          ↓
Statistics (mutex-protected counters)
   Total Fetched, Processed, Errors
```

**All patterns used together:**
- ✅ **Async Fetching**: Fetch from 5 APIs concurrently
- ✅ **Select Statement**: Timeout handling (1 second per fetch)
- ✅ **Worker Pool**: 3 workers process data concurrently
- ✅ **Pipeline**: Data flows through stages (fetch → process → output)
- ✅ **Mutex**: Thread-safe statistics tracking
- ✅ **WaitGroups**: Synchronization at each stage

### Individual Pattern Demonstrations

The code also includes separate demonstrations of each pattern (commented out by default):

#### 1. **Worker Pool Pattern**
- Multiple workers processing jobs concurrently
- Demonstrates goroutines and WaitGroups
- Efficient task distribution across workers
- Results collection through channels

#### 2. **Async Data Fetching**
- Concurrent data fetching from multiple sources
- Parallel I/O operations
- Demonstrates how async operations reduce total execution time
- WaitGroup synchronization

#### 3. **Thread-Safe Counter (Mutex)**
- Demonstrates safe concurrent access to shared resources
- Uses `sync.Mutex` to prevent race conditions
- Shows the importance of synchronization primitives

#### 4. **Pipeline Pattern**
- Multi-stage data processing pipeline
- Channels connecting different processing stages
- Demonstrates data flow through concurrent stages

#### 5. **Select Statement**
- Channel multiplexing
- Handling multiple channel operations
- Timeout handling
- Non-blocking channel operations

## Project Structure

```
Golang/
├── main.go                           # Main program with all concurrency demos
├── go.mod                            # Go module definition
├── run.sh                            # Quick installation & run script
├── README                            # This file
└── Goroutines Tutorial/              # 📚 Complete Goroutines Tutorial
    ├── GETTING_STARTED.md            # Quick start guide
    ├── OVERVIEW.txt                  # Visual overview of tutorial
    ├── QUICK_REFERENCE.txt           # Syntax cheat sheet
    ├── run.sh                        # Interactive menu to run examples
    ├── 01_basic_goroutine.go         # Level 1: Basics
    ├── 02_intermediate_goroutine.go  # Level 2: WaitGroups & Channels  
    ├── 03_advanced_goroutine.go      # Level 3: Advanced Patterns
    ├── exercises.go                  # Practice problems (10 + bonus)
    └── solutions.go                  # Solutions to exercises
```

### 🎓 New to Go? Start Here!

If you're new to Golang and goroutines, check out the **comprehensive tutorial**:

```bash
cd "Goroutines Tutorial"
cat GETTING_STARTED.md   # Quick overview
cat OVERVIEW.txt         # Visual guide
./run.sh                 # Interactive examples
```

The "Goroutines Tutorial" folder contains:
- Step-by-step tutorial from beginner to advanced
- 3 levels of examples with increasing complexity
- 10 practice exercises with solutions
- Quick reference card for syntax
- Interactive menu to run all examples

## Prerequisites

- **Go 1.21 or higher** installed on your system
- Basic understanding of programming concepts

### Installing Go

**On Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install golang-go
```

**Or using snap:**
```bash
sudo snap install go --classic
```

**On macOS:**
```bash
brew install go
```

**On Windows:**
Download from https://golang.org/dl/

Verify installation:
```bash
go version
```

## How to Run

### Method 1: Direct Run
```bash
cd <project root>
go run main.go
```

### Method 2: Build and Execute
```bash
cd <project root>

# Build the binary
go build -o concurrency-demo

# Run the binary
./concurrency-demo
```

### Method 3: Install Globally
```bash
cd <project root>
go install

# Run from anywhere
concurrency-demo
```

## Visual Explanation of Concurrency Patterns

### 1. Worker Pool Pattern

```
                    ┌─────────────┐
                    │  Job Queue  │
                    │  (channel)  │
                    └──────┬──────┘
                           │
           ┌───────────────┼───────────────┐
           │               │               │
           ▼               ▼               ▼
    ┌──────────┐    ┌──────────┐    ┌──────────┐
    │ Worker 1 │    │ Worker 2 │    │ Worker 3 │
    │(goroutine│    │(goroutine│    │(goroutine│
    └─────┬────┘    └─────┬────┘    └─────┬────┘
          │               │               │
          └───────────────┼───────────────┘
                          ▼
                   ┌─────────────┐
                   │   Results   │
                   │  (channel)  │
                   └─────────────┘

How it works:
• Multiple workers (goroutines) pull jobs from shared queue
• Each worker processes jobs independently
• Results sent to results channel
• Efficient for handling many tasks with limited resources
```

### 2. Async Data Fetching (Fan-Out)

```
Sequential (Slow):
API-1 --> [100ms] --> API-2 --> [100ms] --> API-3 --> [100ms]
Total: 300ms

Concurrent (Fast):
           +--- API-1 [100ms] ---+
           |                     |
Main ------+--- API-2 [100ms] ---+---> Collect Results
           |                     |
           +--- API-3 [100ms] ---+
Total: ~100ms (all run in parallel!)

Visualization:
    Main Thread
        |
        +---- go fetch(API-1) ----> result
        |
        +---- go fetch(API-2) ----> result
        |
        +---- go fetch(API-3) ----> result
        |
        +---- WaitGroup.Wait() ---> All Done!
```

### 3. Mutex (Thread-Safe Access)

```
WITHOUT Mutex (Race Condition ❌):

Goroutine 1:  Read counter=0  --> Add 1 --> Write counter=1
Goroutine 2:      Read counter=0  --> Add 1 --> Write counter=1
Result: counter=1 (WRONG! Should be 2)


WITH Mutex (Safe ✅):

Goroutine 1:  Lock --> Read=0 --> Add 1 --> Write=1 --> Unlock
Goroutine 2:        (waiting...)  Lock --> Read=1 --> Add 1 --> Write=2 --> Unlock
Result: counter=2 (CORRECT!)

Visualization:
    ┌─────────────┐
    │   Counter   │ <-- Protected Resource
    │   (shared)  │
    └──────┬──────┘
           |
      ┌────┴────┐
      │  Mutex  │ <-- Lock/Unlock controls access
      └────┬────┘
           |
    ┌──────┴──────────┐
    |                 |
Goroutine 1     Goroutine 2
(waits turn)    (gets lock first)
```

### 4. Pipeline Pattern

```
Stage 1: Generator --> Stage 2: Processor --> Stage 3: Consumer

   ┌──────────┐         ┌──────────┐         ┌──────────┐
   │Generator │ --ch1-->│Processor │ --ch2-->│ Consumer │
   │  (1-10)  │         │ (square) │         │ (print)  │
   └──────────┘         └──────────┘         └──────────┘

Data Flow Example:
Generator: 1 --> Processor: 1² = 1  --> Consumer: "Result: 1"
Generator: 2 --> Processor: 2² = 4  --> Consumer: "Result: 4"
Generator: 3 --> Processor: 3² = 9  --> Consumer: "Result: 9"

Each stage runs concurrently!
- Generator sends numbers immediately
- Processor squares them as they arrive
- Consumer prints results as they're ready
```

### 5. Select Statement (Multiplexing)

```
Multiple channels, one receiver:

        Channel 1 --+
                    |
        Channel 2 --+---> SELECT ---> Handle whichever is ready first
                    |
        Timeout  ---+

Example:
    select {
    case msg := <-ch1:     --> "Got message from channel 1"
    case msg := <-ch2:     --> "Got message from channel 2"
    case <-time.After(1s): --> "Timeout! No message in 1 second"
    default:               --> "No channel ready right now"
    }

Visualization:
         ┌─────────┐
         │  Ch 1   │--------+
         └─────────┘        |
                            v
         ┌─────────┐    ┌────────┐
         │  Ch 2   │--->│ SELECT │---> First ready wins!
         └─────────┘    └────────┘
                            ^
         ┌─────────┐        |
         │ Timeout │--------+
         └─────────┘
```

### How They Work Together

```
Real-World Integrated Example: API Data Processor (main.go)

┌─────────────────────────────────────────────────────────────────┐
│ STAGE 1: Async Fetching (with SELECT for timeout)               │
└─────────────────────────────────────────────────────────────────┘
    
    go fetch(API-1) --+
    go fetch(API-2) --+
    go fetch(API-3) --+---> fetchedData channel
    go fetch(API-4) --+      (or timeout after 1s)
    go fetch(API-5) --+

                    |
                    v

┌─────────────────────────────────────────────────────────────────┐
│ STAGE 2: Worker Pool (3 workers process data)                   │
└─────────────────────────────────────────────────────────────────┘

        ┌──────────────┐
        │ fetchedData  │
        │  (channel)   │
        └──────┬───────┘
               |
    ┌──────────+───────────┐
    |          |           |
    v          v           v
┌────────┐ ┌────────┐ ┌────────┐
│Worker 1│ │Worker 2│ │Worker 3│
└───┬────┘ └───┬────┘ └───┬────┘
    |          |          |
    +----------+----------+
               v
        ┌─────────────┐
        │processedData│
        │  (channel)  │
        └──────┬──────┘
               |
               v
┌─────────────────────────────────────────────────────────────────┐
│ STAGE 3: Pipeline Output                                        │
└─────────────────────────────────────────────────────────────────┘

    for result := range processedData {
        print(result)
    }

                |
                v
┌─────────────────────────────────────────────────────────────────┐
│ STATISTICS (Mutex-Protected)                                    │
└─────────────────────────────────────────────────────────────────┘

    var mu sync.Mutex
    ┌────────────────┐
    │ Fetched: 5     │ <-- Protected by mutex
    │ Processed: 5   │     Multiple goroutines
    │ Errors: 0      │     safely update
    └────────────────┘

Running the program shows:
1. All 5 APIs fetched concurrently (notice varying times)
2. 3 workers process data as it arrives
3. Results printed in pipeline
4. Final statistics displayed
```

## Expected Output

When you run `main.go`, you'll see the **integrated demonstration**:

```
===========================================
  Go Concurrency & Async Programming Demo
===========================================

=== 🎯 INTEGRATED DEMO: All Patterns Combined ===

🌐 Stage 1: Fetching from multiple APIs concurrently...
   ✓ Fetched from API-4 in 289ms
   ✓ Fetched from API-3 in 314ms
   ✓ Fetched from API-1 in 518ms
   ✓ Fetched from API-2 in 606ms
   ✓ Fetched from API-5 in 785ms

⚙️  Stage 2: Processing data with worker pool...

📦 Processing Results:
   Worker-3: PROCESSED[data-from-API-3] (from API-3)
   Worker-2: PROCESSED[data-from-API-4] (from API-4)
   Worker-1: PROCESSED[data-from-API-1] (from API-1)
   Worker-3: PROCESSED[data-from-API-2] (from API-2)
   Worker-2: PROCESSED[data-from-API-5] (from API-5)

📊 Final Statistics:
   Fetched: 5 | Processed: 5 | Errors: 0

✅ Integrated demo completed!

Patterns used:
   ✓ Async Fetching: Concurrent API calls
   ✓ Select: Timeout handling
   ✓ Worker Pool: Limited concurrent processors
   ✓ Pipeline: Data flows through stages
   ✓ Mutex: Thread-safe statistics
   ✓ WaitGroups: Synchronization at each stage
```

**Notice:**
- APIs are fetched concurrently (notice different completion times)
- Workers process data as soon as it arrives
- Results appear in order of processing completion
- No race conditions (safe concurrent access to statistics)

**To see individual pattern demonstrations**, uncomment the respective functions in `main.go`:
```go
// Uncomment to see individual demos:
demonstrateWorkerPool()
demonstrateAsyncFetching()
demonstrateMutex()
demonstratePipeline()
demonstrateSelect()
```

## Code Highlights

### Goroutines
```go
go worker.Process(i, &wg, results)  // Launch a goroutine
```

### Channels
```go
results := make(chan string, 10)    // Buffered channel
dataChan <- data                    // Send to channel
result := <-dataChan                // Receive from channel
```

### WaitGroup
```go
var wg sync.WaitGroup
wg.Add(1)        // Increment counter
defer wg.Done()  // Decrement counter
wg.Wait()        // Block until counter is 0
```

### Mutex
```go
var mu sync.Mutex
mu.Lock()        // Acquire lock
// Critical section
mu.Unlock()      // Release lock
```

### Select Statement
```go
select {
case msg := <-chan1:
    // Handle message from chan1
case msg := <-chan2:
    // Handle message from chan2
case <-time.After(timeout):
    // Handle timeout
}
```

## Learning Resources

- **Official Go Documentation**: https://golang.org/doc/
- **Go by Example**: https://gobyexample.com/
- **A Tour of Go**: https://tour.golang.org/
- **Effective Go**: https://golang.org/doc/effective_go
- **Go Concurrency Patterns**: https://go.dev/blog/pipelines

## Performance Characteristics

- **Goroutines**: ~2KB initial stack size, can grow/shrink dynamically
- **Context Switching**: Much faster than OS thread context switching
- **Scalability**: Can run millions of goroutines on a single machine
- **Scheduler**: Work-stealing scheduler for optimal CPU utilization

## Common Use Cases for Go

- **Microservices**: Fast, efficient, built-in HTTP server
- **Cloud Infrastructure**: Docker, Kubernetes written in Go
- **DevOps Tools**: Terraform, Prometheus, etc.
- **Network Programming**: Excellent concurrency for handling connections
- **CLI Tools**: Fast compilation, single binary distribution
- **Data Processing**: Concurrent pipelines for high throughput

## Tips for Go Concurrency

1. **Start goroutines responsibly**: Always ensure they can complete
2. **Close channels when done**: The sender should close channels
3. **Use WaitGroups**: To wait for multiple goroutines to complete
4. **Avoid sharing memory**: Use channels to communicate
5. **Be careful with closures**: Capture loop variables properly
6. **Use context for cancellation**: For graceful shutdown

## Troubleshooting

### Race Conditions
Run with race detector:
```bash
go run -race main.go
```

### Deadlocks
Go will detect and report deadlocks:
```
fatal error: all goroutines are asleep - deadlock!
```

## Contributing

Feel free to add more concurrency patterns and examples to this project!

## License

This is a demonstration project for educational purposes.

---

**Created for**: Redfish Entry Test Project  
**Purpose**: Demonstrate Go's concurrency features  
**Author**: Educational Demo  
**Last Updated**: December 2025
