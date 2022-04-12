package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)

	target := rand.Intn(100) + 1
	fmt.Println("Random number selected!")

	success := false

	for gusses := 0; gusses < 10; gusses++ {

		fmt.Println("Attempts ", 10-gusses)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess > target {
			fmt.Println("Your number is higher")
		} else if guess < target {
			fmt.Println("Your number is less")
		} else {
			success = true
			fmt.Println("You won")
			break
		}
	}

	if !success {
		fmt.Println("You lose number is ", target)
	}
}
