/*
All the following are basic types available in go
- bool
- string
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr

- byte // alias for uint8

- rune // alias for int32
       // represents a Unicode code point

- float32 float64
- complex64 complex128
*/

package main

import (
  "fmt"
  "math/cmplx"
)

var (
  Tobe bool = false
  MaxInt uint64 = (1 << 64) - 1
  z complex128 = cmplx.Sqrt( -5 + 12i)
)

func main() {
  const f = "%T(%v)\n"  // %T --> type,  %v --> value
  fmt.Printf(f, Tobe, Tobe)
  fmt.Printf(f, MaxInt, MaxInt)
  fmt.Printf(f, z, z)
}
