/**
 *  Detemine square root of any number using newton's method
 *  Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:
 *  z = z - (z * z - x)/(2 * z)
 *  change the loop condition to stop once the value has stopped changing (or only changes by a very small delta).
 *  See if that's more or fewer iterations. How close are you to the math.Sqrt?
 */

 package main

 import (
   "fmt"
   "math"
 )

 func squareRoot(x float64) float64 {
  z := float64(2)
  s := float64(0)
  for {
    z = z - (z * z - x)/(2 * z)
    if ( math.Abs(s -z ) < 1e-15 ) {
      break
    }
    s = z
  }
  return s
 }

 func main() {
   fmt.Println("Square root of 3 is ", squareRoot(3))
 }
