package main

import (
	"fmt"
	"github.com/ptsypyshev/gb-golang-level02/lesson02/fibonacci"
	"os"
)

func main() {
	var stopNumber int
	fmt.Print("Enter your number: ")
	_, err := fmt.Scan(&stopNumber)
	if err != nil {
		fmt.Println("Incorrect input")
		os.Exit(1)
	}
	fmt.Printf("Fibonacci number %d has value %d\n", stopNumber, fibonacci.FibWithCache(stopNumber))
}
