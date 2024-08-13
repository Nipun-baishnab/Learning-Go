package main

import (
	"fmt"
	"os"
)

func readFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("readFile: %w", err)
	}
	return nil
}

func main() {
	if err := readFile("nonexistent.txt"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File opened successfully")
	}
}
