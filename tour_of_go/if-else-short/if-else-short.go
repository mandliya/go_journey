/**
 *  The variable declared inside if short statement are also available in else of that if
 */

package main

import (
  "fmt"
  "math"
)

func pow( x, y, lim float64) float64 {
  if v := math.Pow(x, y); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // v cant be used here though
  return lim
}

func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}
