// Unlike other C type languages, Go can return multiple values from a function
// Here swap will return two values.

package main

import "fmt"

func swap( a, b string ) (string, string ) {
  return b, a
}

func main() {
  a , b := swap("Hello", "World")
  fmt.Println("After swapping string 'hello world', it becomes", a, b )
}
