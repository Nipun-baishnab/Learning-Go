package main

import "fmt"

// Generate a sequence of numbers
func generateNumbers() chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()
	return ch
}

// Double each number
func doubleNumbers(input chan int) chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			output <- num * 2
		}
	}()
	return output
}

// Square each number
func squareNumbers(input chan int) chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			output <- num * num
		}
	}()
	return output
}

// Subtract one from each number
func subtractOne(input chan int) chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			output <- num - 1
		}
	}()
	return output
}

// Print each number
func printNumbers(input chan int) {
	for num := range input {
		fmt.Println(num)
	}
}

func main() {
	// Set up the pipeline
	numbers := generateNumbers()
	doubled := doubleNumbers(numbers)
	squared := squareNumbers(doubled)
	subtracted := subtractOne(squared)
	printNumbers(subtracted)
}
