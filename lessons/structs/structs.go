package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "time"
)

type User struct {
  firstName string
  lastName string
  birthDate string
  createdDate time.Time
}

var reader = bufio.NewReader(os.Stdin)

func main() {
  var newUser User
  firstName := getUserData("Please enter your first name: ")
  lastName := getUserData("Please enter your last name: ")
  birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
  created := time.Now()

  // newUser = User {
  //   firstName: firstName, 
  //   lastName: lastName
  //   birthDate: birthDate,
  //   createdDate: created,
  // }

  // newUser = User {
  //   firstName,
  //   lastName,
  //   birthDate,
  //   created,
  // }

  newUser = User{
    firstName: firstName,
    birthDate: birthDate,
    dateCreated: created,
  }
  
  // ... do something awesome with that gathered data!

  //fmt.Println(firstName, lastName, birthDate, created)
  fmt.Println(newUser)
}

func getUserData(promptText string) string {
  fmt.Print(promptText)
  userInput, _ := reader.ReadString('\n')
  cleanedInput = strings.Replace(userInput, "\n", -1)
  return cleanedInput
}
