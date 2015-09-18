/**
 *  Go also has concept of pointer like c/c++
 *  The type *T is a pointer to a T value. Its zero value is nil.
 *  var p *int
 *  The & operator generates a pointer to its operand.
 *
 *  i := 42
 *  p = &i
 *  The * operator denotes the pointer's underlying value.
 *  fmt.Println(*p) // read i through the pointer p
 *  *p = 21         // set i through the pointer p
 *  This is known as "dereferencing" or "indirecting".
 *  NOTE Unlike C, Go has no pointer arithmetic.
 */

package main

import "fmt"
func main() {
  i, j := 42, 2701
  // point to i
  p := &i
	// read i through the pointer
  fmt.Println(*p)
  // set i through the pointer
  *p = 43
  // see the new value of i
  fmt.Println(i)

  // point to j
  p = &j
	//divide j through the pointer
  *p = *p / 37
  // see the new value of j
  fmt.Println(j)
}
