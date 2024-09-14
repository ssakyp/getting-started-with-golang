package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)
	weight, _:= strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)
	// fmt.Print(weightInput, heightInput)

	//Calculate the BMI  (weight / (height * heigh))
	bmi := weight / (height * height)
	//Output the calculated BMI
	fmt.Printf("Your Body Mass Index is %.2f", bmi)
}
