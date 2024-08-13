package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sayHello(id int, wg *sync.WaitGroup) {
	wg.Done()
	sleepTime := rand.Intn(10000) // Random sleep time up to 1000ms
	fmt.Printf("Goroutine %d will sleep for %d ms\n", id, sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Printf("Hello from Goroutine %d after sleeping for %d ms!\n", id, sleepTime)
}

func main() {

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go sayHello(i, &wg)
	}
	wg.Wait()
	fmt.Println("all goroutines finished")
}
