package main
import "fmt"

func main() {
  age := 32
  fmt.Println(age) //32

  var myAge *int
  myAge = &age
  fmt.Println(myAge) //0xc00001c070
}
