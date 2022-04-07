package main

import "fmt"

func main() {
	var v1 int
	fmt.Println(v1)

	var v2 int = 100
	fmt.Println(v2)

	v3 := 5
	fmt.Println(v3)

	v4 := 5
	v4 = 10
	fmt.Println(v4)

	var v5, v6 = 14, 25
	fmt.Println(v5, v6)

	var (
		v7 = 11
		v8 = "string"
	)
	fmt.Println(v7, v8)
}
