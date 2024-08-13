package main

import "fmt"

func calc(xs []float64) float64 {
	tot := 0.0
	for _, v := range xs {
		tot += v
	}
	return tot / float64(len(xs))
}

func main() {
	xs := []float64{1.0, 2.0, 3.0}
	fmt.Println(calc(xs))
}
