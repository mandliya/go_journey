/**
 * Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers.
 */

package main

import "fmt"

func fibonacci() func() int {
  x := 0
  y := 1
  n := 0
  return func () int{
    switch {
    case n == 0:
      n++
      return 0
    case n == 1:
      n++
      return 1
    default :
      z := x + y
      x = y
      y = z
      return z
    }
  }
}

func main()  {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
