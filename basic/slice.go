package main

import "fmt"

func main() {

	notes := []string{
		"do",
		"re",
		"mi",
	}

	fmt.Println(notes[:1])

	notes = append(notes, "re")

	fmt.Println(notes)

}
