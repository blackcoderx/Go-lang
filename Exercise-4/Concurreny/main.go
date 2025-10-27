package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1. The 'Result' Blueprint
// This defines a structure to hold the result of scraping a single website.
type Result struct {
	URL     string // Which site was scraped?
	Success bool   // Was it successful?
	Words   int    // How many words were found? (if successful)
	Error   string // What was the error? (if it failed)
}

// 2. The 'scrape' Worker
// This function pretends to scrape a single website.
func scrape(url string, ch chan Result, wg *sync.WaitGroup) {
	// Tell the main program we are done when this function finishes.
	defer wg.Done()
	// Pretend it takes 2 seconds to get data from the site.
	time.Sleep(2 * time.Second)

	// Simulate a 30% chance of failure (e.g., website is down).
	if rand.Float64() < 0.3 {
		// If it fails, send a failure result back through the channel.
		ch <- Result{
			URL:     url,
			Success: false,
			Error:   "Connection timeout",
		}
		return // Stop here.
	}

	// If it succeeds, generate a random word count.
	words := rand.Intn(1000) + 100
	// Send a success result back through the channel.
	ch <- Result{
		URL:     url,
		Success: true,
		Words:   words,
	}
}

// 3. The 'main' Function (The Manager)
func main() {
	// A WaitGroup helps the manager wait for all workers to finish.
	var wg sync.WaitGroup
	// A channel is like a conveyor belt for workers to send results back to the manager.
	ch := make(chan Result)

	// The list of websites to check.
	sites := []string{"site1.com", "site2.com", "site3.com", "site4.com", "site5.com"}

	// 4. Start all the workers.
	for _, site := range sites {
		// Tell the WaitGroup we are starting one more worker.
		wg.Add(1)
		// Start a new worker in the background to scrape a site.
		// The 'go' keyword is what makes it run concurrently.
		go scrape(site, ch, &wg)
	}

	// 5. A helper to close the conveyor belt.
	// This runs in the background and waits for all workers to finish.
	// Once they are all done, it closes the channel.
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 6. Collect and process the results.
	successCount := 0
	failureCount := 0
	totalWords := 0

	// Loop over the channel (conveyor belt) and grab each result as it comes in.
	// This loop will automatically end when the channel is closed.
	for result := range ch {
		if result.Success {
			fmt.Printf("✓ %s - %d words\n", result.URL, result.Words)
			successCount++
			totalWords += result.Words
		} else {
			fmt.Printf("✗ %s - FAILED: %s\n", result.URL, result.Error)
			failureCount++
		}
	}

	// 7. Print the final summary.
	fmt.Println("\n--- Summary ---")
	fmt.Printf("Successful: %d\n", successCount)
	fmt.Printf("Failed: %d\n", failureCount)
	fmt.Printf("Total words: %d\n", totalWords)
}
