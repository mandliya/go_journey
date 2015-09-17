//In Go, a name is exported if it begins with a capital letter.
//Foo is an exported name, as is FOO. The name foo is not exported.

package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println("The value of pi is :", math.Pi)    //math.pi won't work, exported names start with a capital letter.
}
