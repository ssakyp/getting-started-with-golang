package main

import "fmt"

func main() {
  // websites := []string{"https://google.com", "https://aws.com"} this would be problem
  websites := map[string]string{
    "Google": "https://google.com",
    "Amazon Web Services": "https://aws.com",
  }
  fmt.Println(websites)
  fmt.Println(websites["Amazon Web Services"])
  websites["LinkedIn"] = "https://linkedin.com" // can be easily added
  delete(websites, "Google")
  fmt.Println(websites) // no google
  
}
