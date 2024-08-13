package main

import (
	"fmt"
	"sync"
)

func printNumbers(id, from, to int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	//sleepTime := rand.Int() % 1000
	//time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	mutex.Lock()
	for i := from; i <= to; i++ {
		fmt.Println(i, "in goroutine", id)
		//time.Sleep(time.Second) // Simulate some work or delay
	}
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	ranges := [][2]int{
		{1, 5},
		{6, 10},
		{11, 15},
	}

	for id, r := range ranges {
		from, to := r[0], r[1]
		wg.Add(1)
		go printNumbers(id+1, from, to, &wg, &mutex)
	}

	wg.Wait()
	fmt.Println("All goroutines have finished.")
}
