// NOTE special way of declaration except in package scope
//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.


// All declarations below are valid way of declaring.
package main

import "fmt"

func main() {
  var i, j int = 1, 2
  k := 3
  var l = 4
  var c, python = false, "No!"
  java := false
  var cpp string = "Nope!"
  fmt.Println(i, j, k , l, c, python, java, cpp)
}
