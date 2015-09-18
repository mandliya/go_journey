/**
 * Struct in Go
 */

package main

import "fmt"

type Vertex struct {
  X, Y int
}

func main() {
  v := Vertex{ 1, 2 }
  fmt.Println("Current Vertex values", v)
  v.X = 4
  fmt.Println("Changed to", v)
}
