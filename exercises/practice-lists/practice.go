package main

import "fmt"

func main() {
	hobbies := [3]string{"tennis", "filmmaking", "coding"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	var firstHobbies []string
	firstHobbies = hobbies[0:2]
	secondHobbies := []string{hobbies[0], hobbies[1]}
	fmt.Println(firstHobbies)
	fmt.Println(secondHobbies)
	firstHobbies = firstHobbies[1:3]
	fmt.Println(firstHobbies)

	dynamicArray := []string{"software engineer", "algorithm expert"}
	fmt.Println(dynamicArray)
	dynamicArray[1] = "AlgoExpert"
	dynamicArray = append(dynamicArray, "systems expert")
	fmt.Println(dynamicArray)
	

}
