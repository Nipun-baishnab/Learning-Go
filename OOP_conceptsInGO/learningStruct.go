package main

import "fmt"

type NamedObj struct {
	Name1, Name2 string
}

type Shape struct {
	NamedObj  //inheritance
	color     int32
	isRegular bool
}

type Point struct {
	x, y float64
}

type Rectangle struct {
	NamedObj            //multiple inheritance
	Shape               //^^
	center        Point //standard composition
	Width, Height float64
}

func main() {
	var aRect = Rectangle{
		NamedObj{"A Random Rectangle", "A non Square Actually"},
		Shape{NamedObj{"5*7 Rectangle", "5*7 Reall?"}, 574, true},
		Point{0.0, 0.0},
		20.0, 2.5,
	}
	fmt.Println(aRect.Name1, aRect.Name2)
	fmt.Println(aRect.Shape)
	fmt.Println(aRect.Shape.Name1)
}
