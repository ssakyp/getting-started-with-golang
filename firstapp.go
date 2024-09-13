package main

import "fmt"

func main() {
  //fmt.Println("Hello World!")
  //var greetingText string
  //greetingText = "Hello World!"
  //var greetingText string = "Hello World!"
  greetingText := "Hello World!"
  fmt.Println(greetingText)
  
  luckyNumber := 17
  evenMoreLuckyNumber := luckyNumber + 5
  fmt.Println(luckyNumber) //17
  fmt.Println(evenMoreLuckyNumber) //22

  var newNumber float64 = float64(luckyNumber) / 3
  fmt.Println(newNumber) // 5...

  var defaultFloat float64 = 1.123456789145623123456
  var smallFloat float32 = 1.123456789145623123456

  fmt.Println(defaultFloat) //1.1234567891456231
  fmt.Println(smallFloat) //1.1234567

  var firstRune rune = '€'
  fmt.Println(string(firstRune)) //8364
  fmt.Println(firstRune) //€

  var firstByte byte = 'a'
  fmt.Println(firstByte) // 97
  fmt.Println(string(firstByte)) // a

  firstName := "Sultan"
  lastName := "Sakyp"
  fullName := firstName + " " + lastName

  fmt.Println("Hi, my full name is " + fullName)
  //fmt.Println("9" + 1) not allowed
  
}
