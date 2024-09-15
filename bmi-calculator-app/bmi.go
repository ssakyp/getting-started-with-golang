package main

import (
	"fmt"

	"github.com/ssakyp/bmi-calculator-app/info"
)

func main() {
	//Output information
	info.PrintWelcome()
	
	weight, height := getUserMetrics()
	//Calculate the BMI  (weight / (height * heigh))
	bmi := weight / (height * height)
	//Output the calculated BMI
	fmt.Printf("Your Body Mass Index is %.2f\n", bmi)
}
