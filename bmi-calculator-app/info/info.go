package info

import "fmt"

const mainTitle = "BMI Calculator"
const separator = "------------------"
const WeightPrompt = "Please enter your weight in (kg):"
const HeightPrompt = "Please enter your height in (m): "

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
}
