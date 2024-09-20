package main

import (
  "os"
  "bufio"
  "fmt"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Please enter your age: ")
  userAgeInput := reader.ReadString('\n')
  userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
  userAge, _ := strconv.ParseInt(userAgeInput, 0, 64)

  // isOldEnough := userAge >= 18
  
  if userAge >= 18 {
    fmt.Println("Welcome to the club!")
  } else {
    fmt.Println("You're not old enough!")
  }

  fmt.Println("Goodbye!")
}
