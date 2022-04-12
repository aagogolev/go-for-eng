package main

import "fmt"

func main() {

	notes := [7]string{
		"do",
		"re",
		"mi",
	}

	fmt.Println(notes)

	for i := 0; i <= 2; i++ {
		fmt.Println(notes[i])
	}

	for index, value := range notes {
		fmt.Println(index, value)
	}

	for _, value := range notes {
		fmt.Println(value)
	}
}
