package main

import "fmt"

func main() {
	a := 5
	b := 6
	sum := add(a, b)
	printNumber(sum)
}

func add(num1, num2 int) int {
	return num1 + num2
}

func printNumber(num int) {
	fmt.Println(num)
}