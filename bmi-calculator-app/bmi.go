package main

import (
	"fmt"
	"strconv"
	"strings"
)

const mainTitle = "BMI Calculator"
const separator = "------------------"
const weightPrompt = "Please enter your weight in (kg):"
const heightPrompt = "Please enter your height in (m): "

func main() {
	//Output information
	fmt.Println(mainTitle)
	fmt.Println(separator)
	//Prompt for user input (weight + height)
	fmt.Println(weightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Println(heightPrompt)
	heightInput, _ := reader.ReadString('\n')

	//Save that user input in variables
	weightInput = strings.Replace(weightInput, "\n", "", -1)
	heightInput = strings.Replace(heightInput, "\n", "", -1)
	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)
	// fmt.Print(weightInput, heightInput)

	//Calculate the BMI  (weight / (height * heigh))
	bmi := weight / (height * height)
	//Output the calculated BMI
	fmt.Printf("Your Body Mass Index is %.2f", bmi)
}
