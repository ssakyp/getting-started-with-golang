package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := 5
	b := 6
	sum := add(a, b)
	printNumber(sum)
	printNumber(generateRandomNumbers())
}

func add(num1, num2 int) int {
	return num1 + num2
}

func printNumber(num int) {
	fmt.Printf("The number is %v\n",num)
}

func generateRandomNumbers() int {
	return rand.Intn(65)
}