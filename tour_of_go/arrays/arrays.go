/**
 * The type [n]T is an array of n values of type T.
 * The expression
 * var a [10]int
 * declares a variable a as an array of ten integers.
 * An array's length is part of its type, so arrays cannot be resized.
 * This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
 */

package main

import "fmt"

func main() {
  var arr [2]string
  arr[0] = "Awesome"
  arr[1] = "GO!"
  fmt.Println(arr[0], arr[1])
  fmt.Println(arr)
}
