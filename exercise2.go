package main

import "fmt"

func main() {
  var pi float64 = 3.14
  var radii int = 5
  var circ float64 = 2 * pi * float64(radii)

  fmt.Println(circ)
  fmt.Printf("For a radius of %v, the circle circumference is %v.", radii, circ)

}
