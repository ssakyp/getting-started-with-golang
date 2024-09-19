package main

import "fmt"

func main() {
  prices := []float64{10.99, 8.99}
  fmt.Println(prices[1]) //8.99
  fmt.Println(prices[0:1]) // 10.99
  prices[1] = 1.11
  //prices[2] => out of range => panics
  updatedPrices := append(prices, 5.99)
  prices = append(prices, 4.87)
  fmt.Println(updatedPrices) // [10.99 1.11 5.99]
  fmt.Println(prices) // [10.99 1.11 4.87]
  
}
