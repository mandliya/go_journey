/**
 * The zero value of a slice is nil.
 * A nil slice has a length and capacity of 0.
 */

package main

import "fmt"

func main() {
  var arr []int
  fmt.Println(arr, len(arr), cap(arr))
  if arr == nil {
    fmt.Println("nil!!")
  }
}
