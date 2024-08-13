package main

import (
	"fmt"
)

func main() {
	elements := make(map[string]string)

	for i := 0; i < 2; i++ {
		var x [2]string
		fmt.Scanf("%s", &x[0])
		fmt.Scanf("%s", &x[1])
		elements[x[0]] = x[1]
	}

	fmt.Println(elements)

}
