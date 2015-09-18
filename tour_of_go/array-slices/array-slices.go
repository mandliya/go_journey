/**
 * A slice points to an array of values and also includes a length.
 * []T is a slice with elements of type T.
 */

package main

import "fmt"

func main() {
  s := []int{2, 3, 5, 7, 11, 13}
  fmt.Println("s ==", s)

  for i := 0 ; i < len(s); i++ {
    fmt.Printf("s[%d] = %d\n", i , s[i])
  }
}
