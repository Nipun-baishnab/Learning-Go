package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello, Goroutine!")
}

func main() {
	go sayHello() // start a new goroutine
	//time.Sleep(time.Second) // Give the goroutine time to execute,
	// commenting will make the main function finish before the gproutine executes.
}
