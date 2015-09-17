/**
 *  For loop is same as C but no surrounding ( ), not even optional
 *  no other kind of loop exists. just for loop exists in go
 *  Also, here only post increment works, may be we will come to know of reason later.
 */

package main

import "fmt"
func main() {
  sum := 0
  for i := 0 ; i < 10 ; i++ {
    sum += i
  }
  fmt.Println("Sum of 0 to 9 is", sum)

  // We can do something like this as we do for while loop in other langauges
  total := 0
  i := 9
  for ; i > 0 ; {
    total += i
    i -= 1
  }
  fmt.Println("Sum of 0 to 9 is", total)

  // infinite loop
  // for {
  // }
}
