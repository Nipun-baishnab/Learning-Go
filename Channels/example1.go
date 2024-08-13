package main

import (
	"fmt"
	"sync"
)

func writeToChannel(ch chan string) {
	mutex := sync.Mutex{}
	mutex.Lock()
	fmt.Printf("Called\n")
	ch <- "Hello, Channel!" // send a value into the channel
	close(ch)
	fmt.Printf("\ni'm here \n")
	mutex.Unlock()
}

func main() {
	ch := make(chan string) // create a channel
	go writeToChannel(ch)   // start a goroutine to write to channel
	msg := <-ch             // Receive a value from the channel
	fmt.Println(<-ch)
}
