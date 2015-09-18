/*
Struct fields can be accessed through a struct pointer.
The indirection through the pointer is transparent.
 */

package main

import "fmt"

type Vertex struct {
  X, Y int
}

func main() {
  v := Vertex{ 1, 2 }
  p := &v
  fmt.Printf("Current values of X and Y are %v and %v\n", p.X, p.Y)
}
