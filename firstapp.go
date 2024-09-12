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

  evenMoreLuckyNumber = luckyNumber * 3
  fmt.Println(evenMoreLuckyNumber) //51
}
