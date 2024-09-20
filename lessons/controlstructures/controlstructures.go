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
  userAge, err := strconv.ParseInt(userAgeInput, 0, 64)

  fmt.Println(err)
  // isOldEnough := userAge >= 18
  
  if (userAge >= 30 && userAge < 50) || userAge >= 50 {
    fmt.Println("You are eligible for our senior jobs.")
  } else if userAge >= 50 {
    fmt.Println("The best age!")
  } else if userAge == 18 {
    fmt.Println("Welcome to the club!")
  } else {
    fmt.Println("Sorry, you're not old enough.")
  }

  fmt.Println("Goodbye!")
}
