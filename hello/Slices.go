package main

import (
	"fmt"
)

type Dog struct {
	Name string
	Age  int
}

func main() {
	jackie := Dog{
		Name: "Jackie",
		Age:  19,
	}
	fmt.Println(jackie)
	fmt.Printf("Jackie Address : %p\n", &jackie)

	samson := Dog{
		Name: "Samson",
		Age:  29,
	}
	fmt.Println(samson)
	fmt.Printf("Samson Address : %p\n", &samson)

	//dogs := []Dog{jackie, samson}

	fmt.Println("")

	fmt.Printf("Lets use dot %s\n", samson.Name)

}
