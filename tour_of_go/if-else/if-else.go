/*
  Understanding how if else blocks work in go
 */

package main

import (
  "fmt"
  "math"
)

func squareRoot( num float64 ) string  {
  if num < 0 {
    return squareRoot(-num) + "i"
  } else {
    return fmt.Sprint(math.Sqrt(num))
  }
}

func main() {
  x := fmt.Sprint("The square root of 2 is : %s ", squareRoot(2))
  fmt.Println(x)
  y := fmt.Sprint("The square root of -4 is : %s", squareRoot(-4))
  fmt.Println(y)
}
