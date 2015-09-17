package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  fmt.Println("A random number between 0 and 10 (including):", rand.Intn(11))
}
