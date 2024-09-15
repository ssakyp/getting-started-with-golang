package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := generateRandomNumbers()
	sum := add(a, b)
	printNumber(sum)
	
}

func add(num1, num2 int) (sum int) {
	sum = num1 + num2
	return sum
}

func printNumber(num int) {
	fmt.Printf("The number is %v\n",num)
}

func generateRandomNumbers() (int, int) {
	return rand.Intn(65), rand.Intn(30)
}