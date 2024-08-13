package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(id, from, to int, wg *sync.WaitGroup, goroutineWg *sync.WaitGroup) {
	defer wg.Done()
	defer goroutineWg.Done() // Decrease the goroutine's WaitGroup when finished

	for i := from; i <= to; i++ {
		fmt.Println(i, " in goroutine ", id)
		time.Sleep(time.Second) // Simulate some work or delay
	}
}

func main() {
	var wg sync.WaitGroup // Main WaitGroup to wait for all initial goroutines to start

	// Define ranges for each goroutine
	ranges := [][2]int{
		{1, 5},
		{6, 10},
		{11, 15},
	}

	// Iterate over ranges and start a goroutine for each range
	for id, r := range ranges {
		from, to := r[0], r[1]

		wg.Add(1)                        // Increment main WaitGroup for each goroutine
		goroutineWg := &sync.WaitGroup{} // Create a separate WaitGroup for each goroutine
		goroutineWg.Add(1)               // Increment goroutine's own WaitGroup

		go printNumbers(id+1, from, to, &wg, goroutineWg)

		// Wait for the goroutine to finish its work
		go func(wg *sync.WaitGroup) {
			goroutineWg.Wait()
			wg.Done() // Decrease main WaitGroup when the goroutine completes
		}(&wg)
	}

	wg.Wait() // Wait for all initial goroutines to start
	fmt.Println("All goroutines have finished.")
}
