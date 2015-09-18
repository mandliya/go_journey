/**
 * Slices can be re-sliced, creating a new slice value that points to the same array.
 * The expression
 * s[lo:hi]
 * evaluates to a slice of the elements from lo through hi-1, inclusive. Thus
 * s[lo:lo] is empty and
 * s[lo:lo+1] has one element.
 */

package main

import "fmt"

func main() {
  s := []int{ 2, 3, 5, 7, 11, 13}
  fmt.Println("s ==", s)
  fmt.Println("s[1:4] ==", s[1:4])

  // missing low index value applies from 0
  fmt.Println("s[:3] ==", s[:3])

  //missing high index value applies len(s)
  fmt.Println("s[3:] ==", s[3:])
}
