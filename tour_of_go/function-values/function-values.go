/**
 *  Function are values
 */

package main

import (
  "math"
  "fmt"
)

func main() {
  hypot := func(x,y float64) float64 {
    return math.Sqrt(x * x + y * y)
  }
  fmt.Println("The hypotenuse of triangle with sides 3 and 4 is ",hypot(3, 4))
}
