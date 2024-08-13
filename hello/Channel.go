package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
	time.Sleep(10 * time.Second)
	ch <- 42 // Send a value to the channel
}

func main() {
	ch := make(chan int)

	go worker(ch) // Start a goroutine that will send a value to the channel

	fmt.Println("Waiting for value from channel...")
	value := <-ch // Receive the value from the channel
	fmt.Println("Received value:", value)
}
