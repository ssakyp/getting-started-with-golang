package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "time"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
  firstName := getUserData("Please enter your first name: ")
  lastName := getUserData("Please enter your last name: ")
  birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
  created := time.Now()

  // ... do something awesome with that gathered data!

  fmt.Println(firstName, lastName, birthdate, created)
}

func getUserData(promptText string) string {
  fmt.Print(promptText)
  userInput, _ := reader.ReadString('\n')
  cleanedInput = strings.Replace(userInput, "\n", -1)
  return cleanedInput
}
