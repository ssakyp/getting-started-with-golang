package main
import "fmt"

func main() {
  age := 32
  fmt.Println(age) //32

  //var myAge *int
  
  myAge := &age
  *myAge = 33 / changing the value
  
  fmt.Println(myAge) //0xc00001c070
  fmt.Println(*myAge) // 33 => dereferencing prints the value stored on that address

  doubledAge := double(myAge)
  fmt.Println(doubledAge) // 66
  fmt.Println(age) //66
}

//go is creating a copy

func double(number *int) int {
  result := *number * 2
  return result
}
