// Variables in go are declared with var statement, check this example.
//  the type is last. Here its bool.
//A var statement can be at package or function level. We see both in this example.
// var not initialized explicitly will have their zero value
// example int will have 0, boolear will have false, empty string "" etc
package main

import "fmt"

var python, c bool

func main() {
  var i int
  var java string
  fmt.Println(i, java, python, c) // output would be--> 0   false false (extra space is java representing "")   
}
