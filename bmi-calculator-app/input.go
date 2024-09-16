package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ssakyp/bmi-calculator-app/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weight float64, height float64) {
	weight = getUserInput(info.WeightPrompt)
	height = getUserInput(info.HeightPrompt)
	return 
}

func getUserInput(promptText string) (enteredValue float64) {
	fmt.Println(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	enteredValue, _ = strconv.ParseFloat(userInput, 64)
	return
}
