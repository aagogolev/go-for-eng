package main

import (
	"fmt"
	"go-for-eng/basic/calendar"
	"log"
)

func main() {

	date := calendar.Date{}
	err := date.SetYear(3000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)

	fmt.Println(date.Day())

}
