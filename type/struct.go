package main

import "fmt"

func main() {
	type Point struct {
		X int
		Y int
	}

	p := Point{
		X: 5,
		Y: 19,
	}

	p = Point{5, 15}

	fmt.Println(p.X)
	fmt.Println(p.Y)

	p.X = 55
	fmt.Println(p.X)

}
