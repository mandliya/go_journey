/**
 * A struct literal denotes a newly allocated struct value by listing the values of its fields.
 * You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
 * The special prefix & returns a pointer to the struct value.
 */

package main

import "fmt"

type Vertex struct {
  X, Y int
}

func main() {
  v1 := Vertex{ 1, 2 }
  v2 := Vertex{ X: 1 }  //Y : 0 implicitly
  v3 := Vertex{ } // X: 0 , Y : 0 implicitly
  p := &Vertex{ 4, 5 }

  fmt.Println(v1, v2, v3, p)
}
