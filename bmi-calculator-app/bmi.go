package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	//Output information
	fmt.Println("BMI Calculator")
	fmt.Println("------------------")
	//Prompt for user input (weight + height)
	fmt.Println("Please inter your weight in (kg):")
	weightInput, _ := reader.ReadString('\n')

	fmt.Println("Please inter your height in (m):")
	heightInput, _ := reader.ReadString('\n')
	//Save that user input in variables
	//Calculate the BMI  (weight / (height * heigh))
	//Output the calculated BMI
}
