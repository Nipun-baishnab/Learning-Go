package main

import (
	"flag"
	"fmt"
)

func main() {

	// Define flags
	name := flag.String("name", "World", "a name to  say hello to")
	age := flag.Int("age", 25, "a age to say hello to")
	verbose := flag.Bool("verbose", false, "verbose output")

	// Parse the output

	flag.Parse()

	if *verbose {
		fmt.Println("Verbose mode enabled")
	}
	fmt.Printf("Hello, %s!\n", *name)
	if *age > 0 {
		fmt.Printf("You are %d years old.\n", *age)
	}
}
