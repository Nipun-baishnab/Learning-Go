package main

import "fmt"

func main() {

	/// arrays

	var x [5]float64
	for i := 0; i < 5; i++ {
		fmt.Scanf("%f", &x[i])
	}

	a, b, c, d, e := x[0:5]

	fmt.Println(a, b, c, d, e)
	/*
		p := 0.0

		for i := 0; i < 5; i++ {
			p += x[i]



		}

		fmt.Println(float64(p) / 5)
		**/
}
