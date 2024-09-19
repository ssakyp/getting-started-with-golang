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

  // Why maps?
  // For maps anything can be used as a key => more flexibility
  // Structs - predefined data structures => key value pairs cannot be deleted
  // Structs - clearly defined list of fields
  // Maps = arrays but instead of indexes keys
  
}
