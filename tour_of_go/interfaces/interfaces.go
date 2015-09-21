/**
 * An interface type is defined by a set of methods.
 * A value of interface type can hold any value that implements those methods.
 */

package main

import (
  "math"
  "fmt"
)

type Vertex struct {
  X, Y float64
}

type MyFloat float64

type Abser interface {
  Abs() float64
}

func (v * Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

func main()  {
  var a Abser
  f := MyFloat(math.Pi)
  v := Vertex{ 3, 4 }

  a = f     //MyFloat implements Abs
  fmt.Println(a.Abs())

  a = &v    //Vertex * implements Abs
  fmt.Println(a.Abs())

  // This would be an error since Vertex * defines Abs but not Vertex.
  //a = v
  //fmt.Println(a.Abs())
}
