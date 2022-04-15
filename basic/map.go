package main

import "fmt"

func main() {

	emloy := map[string]float64{
		"admin": 22,
		"user":  33,
	}

	fmt.Println(emloy)
	fmt.Println(emloy["admin"])

	cars := make(map[string]int)
	cars["bmw"]++
	cars["mers"] = 2000
	fmt.Println(cars)
	cars["mers"] = 9999
	fmt.Println(cars)

	langs := map[string]float64{
		"java":   100,
		"golang": 200,
		"python": 300,
	}

	for key, value := range langs {
		fmt.Println(key, " ", value)
	}
}
