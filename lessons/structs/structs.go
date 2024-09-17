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

func (user *User) outputDetails() {
  fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastName, user.birthDate)
}

func NewUser(fName string, lName string, bDate string) *User {
  created := time.Now()
  user := User{
    firstName: fName,
    lastName: lName,
    createdDate: created,
  }
  return &user
}

var reader = bufio.NewReader(os.Stdin)

func main() {
  var newUser *User
  firstName := getUserData("Please enter your first name: ")
  lastName := getUserData("Please enter your last name: ")
  birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
  // created := time.Now()

  newUser = NewUser(firstName, lastName, birthDate)
  
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

  // newUser = User{
  //   firstName: firstName,
  //   birthDate: birthDate,
  //   dateCreated: created,
  // }
  
  // ... do something awesome with that gathered data!

  //fmt.Println(firstName, lastName, birthDate, created)
  // fmt.Println(newUser)
  // outputUserData(newUser)
  newUser.outputDetails()
}

// func outputUserData(user *User) {
//   fmt.Printf("My name is %v %v (born on %v)", (*user).firstName, user.lastName, user.birthDate)
// }

func getUserData(promptText string) string {
  fmt.Print(promptText)
  userInput, _ := reader.ReadString('\n')
  cleanedInput = strings.Replace(userInput, "\n", -1)
  return cleanedInput
}
