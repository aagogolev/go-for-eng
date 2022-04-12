package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter your age: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	old, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}

	if old > 18 {
		fmt.Println("Hello")
	} else {
		fmt.Println("Exit")
	}

}
