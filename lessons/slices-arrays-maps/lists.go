package main

import "fmt"

// struct => concept allows to group data together into one data package
type Product struct {
  title string
  id string
  price float64
}

// similar pieces of data which needs to be grouped
// many values describing the same thing


//using struct for such things is annoying
// type TemperatureData struct {
//   day1 float64
//   day2 float64
//   day4 float64
//   ...
// }


func main() {
  //var productNames [5]string
  //productNames = [5]string{"A book"}
  var productNames [5]string = [5]string{"A book"}
  prices := [4]float64{10.99, 93.11, 45.11, 20.45}
  fmt.Println(prices)

  productNames[2] = "A carpet"
  
  fmt.Println(productNames) //[A book  A carpet  ]
  fmt.Println(prices[2]) //45.11

  featuredPrices := prices[1:3] //slice => start at index 1 and excluding index 3
  fmt.Println(featuredPrices) //[93.11 45.11]

  secondFeaturedPrices := prices[:3] // slice start from the beginning up to 3 but excluding it
  fmt.Println(secondFeaturedPrices) // [10.99 93.11 45.11]
}








