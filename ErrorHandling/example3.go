package main

import (
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func doSomething() error {
	return &MyError{Code: 404, Message: "Resource not found"}
}

func main() {
	if err := doSomething(); err != nil {
		fmt.Println("Error:", err)
	}
}
