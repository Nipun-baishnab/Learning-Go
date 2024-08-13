package main

import (
	"fmt"
	"sync"
)

func printNumbersSmallLetters(wg *sync.WaitGroup) {
	defer wg.Done() // signal that this goroutine is done
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		//time.Sleep(500 * time.Millisecond)
	}
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c\n", i)
		//time.Sleep(500 * time.Millisecond)
	}
}

func printCapitalLetters(wg *sync.WaitGroup) {
	defer wg.Done() // signal that this goroutine is done
	for i := 'A'; i <= 'H'; i++ {
		fmt.Printf("%c\n", i)
		//time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // we have two goroutines to wait for

	go printNumbersSmallLetters(&wg)
	//go printLetters(&wg)
	go printCapitalLetters(&wg)

	wg.Wait() // wait for all goroutines to finish
}
