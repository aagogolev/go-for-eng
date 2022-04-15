package main

import "fmt"

type subscriber struct {
	name       string
	rate       float64
	active     bool
	homeAddres adress
}

type adress struct {
	street   string
	city     string
	building string
}

func printInfo(s subscriber) {
	fmt.Println(s.name)
	fmt.Println(s.rate)
	fmt.Println(s.active)
	fmt.Println(s.homeAddres)
}

func defaultSubscripber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 100
	s.active = true

	return s
}

func applyDiscount(s *subscriber) {
	s.rate = 1.99
}

func main() {

	adress := adress{
		street:   "pervay",
		city:     "perlof",
		building: "1",
	}

	subscriber1 := defaultSubscripber("admin")
	printInfo(subscriber1)
	applyDiscount(&subscriber1)
	printInfo(subscriber1)
	subscriber1.homeAddres = adress
	printInfo(subscriber1)
}
