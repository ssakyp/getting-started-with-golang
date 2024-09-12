package main

import "fmt"

func main() {
  var firstName string = "Sultan"
  lastName := "Sakyp"
  fmt.Println(firstName," ", lastName)

  currentYear := 2024
  var birthYear = 1996
  age := currentYear - birthYear
  currentYear = currentYear + 1
  fmt.Println(age)
  age = currentYear - birthYear
  fmt.Println(age)

}
