/**
 *  Constants in go are declared using  'const' keyword. := notation is not allowed for const
 *  const can be boolean, numeric and string.
 *  Numeric constants are high precision values.
 */

package main

import "fmt"

const answer = 42

const (
  Big = 1 << 100
  Small = Big >> 99
)

func needInt(x int) int {
  return x * 10 + 1
}

func needFloat(x float64) float64 {
  return x * 0.1
}

func main() {
  const World = "Ravi"
  fmt.Println("Hello", World)
  fmt.Println("Answer to the ultimate question of Life, the Universe, and Everything is", answer)
  const Truth = true
  fmt.Println("Go is Awesome?", Truth)

  fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
  //fmt.Println(needInt(Big)) This will cause overflow
}
