/**
 *  Understanding short if with a short statement,
 *  if statement can start with a short statement to execute before the condition.
 *  Variables declared by the statement are only in scope until the end of the if.
 *  Understand and experiment with scope of v here/
 */

package main

import (
  "fmt"
  "math"
)

func powLimit( x, y, lim float64) float64 {
  if v := math.Pow(x, y); v < lim {
    return v;
  }
  return lim
}

func main() {
  fmt.Println("The value of 3 to power of 2 if limited under 10 is ", powLimit(3,2,10))
  fmt.Println("The value of 3 to power of 3 if limited under 10 is ", powLimit(3,3,10))
}
