package main

import (
	"fmt"
	"os"
)

func readFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := readFile("/home/appscodepc/go/src/ErrorHandling/test.txt"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File opened successfully")
	}
}
