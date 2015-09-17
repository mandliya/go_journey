//Nextafter returns the next representable float64 value after x towards y.

package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Printf("Nextafter( 2,3 ) will be  %g\n", math.Nextafter(2,3))
}
