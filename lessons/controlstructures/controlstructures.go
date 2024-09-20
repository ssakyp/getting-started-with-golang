package main

import (
  "os"
  "bufio"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  reader.ReadString('\n')
}
