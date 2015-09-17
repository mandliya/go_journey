// In go, functions can take one or more arguments.
// NOTE: The type comes after the variable name.
// GO makes an interesting argument for this
// Read the following article for why go chose this way.
// http://blog.golang.org/gos-declaration-syntax

package main

import (
  "fmt"
)

func add( x int, y int ) int {
  return x + y
}

func main() {
  fmt.Println(add( 3, 9 ))
}
