// function - 2
// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.

package main

import (
  "fmt"
)

func multiply( a, b int ) int {
  return a * b
}

func main() {
  fmt.Println("Multiplication of 3 and 9:", multiply(3, 9))
}
