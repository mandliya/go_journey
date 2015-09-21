/**
 * Methods can be associated with a named type or a pointer to a named type.
 *
 * NOTE: Notice here that Scale changes v, and Abs just reads it.
 * What happens when Vertex as receiver instead of Vertex pointer.
 *
 * Scale would do nothing, as it will have its own copy of V and it will mutate it,
 * Which won't reflect on the v which we provided.
 * But Abs would act the same way, as it does not mutate v just reads it.
 */

package main

import (
  "math"
  "fmt"
)

type Vertex struct {
  X, Y float64
}

func (v * Vertex) Scale( f float64 )  {
  v.X = v.X * f
  v.Y = v.Y * f
}

func (v * Vertex) Abs() float64 {
  return math.Sqrt( v.X * v.X + v.Y * v.Y)
}

func main() {
  v := &Vertex{ 3, 4 }
  fmt.Printf("Before Scaling: %+v, Abs: %v\n", v, v.Abs())
  v.Scale(5)
  fmt.Printf("After Scaling:  %+v, Abs: %v\n", v, v.Abs())
}
