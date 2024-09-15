package main

import (

	"github.com/ssakyp/bmi-calculator-app/info"
)

func main() {
	//Output information
	info.PrintWelcome()
	
	weight, height := getUserMetrics()
	//Calculate the BMI  (weight / (height * heigh))
	bmi := calculateBMI(weight, height)
	printBMI(bmi)
}

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}