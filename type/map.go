package main

import "fmt"

func main() {
	var m1 map[int32]bool
	var m2 map[string]string

	m3 := make(map[string]int)

	ages := map[string]int{
		"admin": 30,
		"user1": 14,
	}

	age := ages["admin"]
	ages["user2"] = 16
	fmt.Println(ages["user3"])
	ages["user3"]++

	fmt.Println(m1, m2, m3, age, ages)

}
