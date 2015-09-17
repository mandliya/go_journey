// A var declaration can include initializers, one per variable.
//If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

package main

import "fmt"
//both delcarations here are valid
//Notice absence of type in second declaration
// c is boolean, while java is string
var integer1, integer2 int = 5, 6

func main() {
  var c, java = false, "No its not java!"
  fmt.Println(integer1, integer2, c, java)
}
