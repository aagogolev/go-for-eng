package main

import "fmt"

func main() {

	sl := []int64{9, 8, 7}
	for key, value := range sl {
		fmt.Println("key: %v, val: %v \n", key, value)
	}

	for _, value := range sl {
		fmt.Println(value)
	}

	ages := map[string]int{
		"admin": 50,
		"user":  44,
	}

	for key, value := range ages {
		fmt.Println(key)
		fmt.Println(value)
	}

	word := "simbir"

	for i := 0; i < len(word); i++ {
		fmt.Println(word[i])
		fmt.Println("%1", word[i])
	}

	for key, value := range word {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Println("%1", value)
	}

}
