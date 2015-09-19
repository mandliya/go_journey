/**
 *  To add more elments to go array we use append.
 *  append( []type, a, b, ...)
 *  append([]type, anotherSlice)
 */

package main

import "fmt"

func main() {

  //nil array
  var a []int

  //append works on nil slice
  a = append(a, 0);
  printSlice("a", a)
  //the slice grows as needed
  a = append(a, 1);
  printSlice("a", a)
  //we can add more than one element at a time
  a = append(a, 2, 3, 4, 5)
  printSlice("a", a)

  // NOTE notice the .. here in append
  b := []int{6, 7, 8}
  a = append(a, b...)
  printSlice("a", a)

}

func printSlice( s string, x []int)  {
  fmt.Printf("%s len=%d cap=%d %v\n",
             s, len(x), cap(x), x)
}
