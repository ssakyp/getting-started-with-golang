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

  anotherPrices := prices[1:] // starts from index 1 up to the end of the array
  anotherPrice[0] = 199.99 // this slice is just a window and if something is changed then it affects the original array
  // this slice is just a tiny refernce to that array
  // it is a memory efficient way selecting parts of arrays
  fmt.Println(anotherPrices) // [199.99 45.11 20.45]
  fmt.Println(prices) // [10.99 199.99 45.11 20.45]


  // go saves a metadata for a slice
  // for every slice we have a length and a capacity
  fmt.Println(len(anotherPrices), cap(anotherPrices)) // we can output the length with a built-in function len => 3 | cap is also a default go function => 3
  // the length gives us the number of items in a slice or array
  // the capacity is a bit more complex, it is the same value here, but it would have been different
 
  highlightedPrices := anotherPrices[:1] // every element starting from the index 0 up to index 1 excluding
  fmt.Println(highlightedPrices) // [199.99]
  fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // 1 3

  highlightedPrices = highlightedPrices[:3] // reslicing
  fmt.Println(highlightedPrices) // [199.99 45.11 20.45]
  fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // 3 3

  //we can select more towards right side of the array not the left side
  
}








