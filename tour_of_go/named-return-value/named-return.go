// A go function can return named return value, just like variables.
// These returned variables can be used directly.
// A return statement without arguments returns the current values of the results.
//  This is known as a "naked" return.
//  Naked return statements should be used only in short functions,
//  as with the example shown here. They can harm readability in longer functions.

package main

import "fmt"

func division( dividend, divisor int ) (quotient, remainder int) {
  quotient = dividend / divisor
  remainder = dividend % divisor
  return
}

func main() {
  var dividend, divisor int = 17, 3
  quotient, remainder := division( dividend, divisor )
  s := fmt.Sprintf("The quotient and remainder in division of %d with %d is %d and %d.",
               dividend, divisor, quotient, remainder)
  fmt.Println(s)
}
