package main

import (
  "os"
  "bufio"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Please enter your age: ")
  userAgeInput := reader.ReadString('\n')
  userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
  userAge, _ := strconv.ParseInt(userAgeInput, 0, 64)

  if userAge >= 18 {
    fmt.Println("Welcome to the club!")
  }
}
