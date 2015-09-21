/**
 *  Implement Sqrt function using Newton's method, return an error if passed a negative number
 */

package main

import (
  "fmt"
  "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error)  {
  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }
  z := float64(2.0)
  s := float64(0)
  for {
    z = z - ((z * z) - x)/(2 * z)
    if math.Abs(s -z) < 1e-15 {
      break
    }
    s = z
  }
  return s, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
