/**
 *  One of the most common interface is Stringer
 *  type Stringer interface {
 *  	String() string
 *  	}
 *  	A Stringer is a type that can describe itself as a string.
 *  	The fmt package (and many others) look for this interface to print values.
 */

package main

import "fmt"

type Person struct {
  Name string
  Age  int
}

func (p Person) String() string {
  return fmt.Sprintf("Name : %v, Age : %v\n", p.Name, p.Age)
}

func main() {
  n := Person{ "Narendra Modi", 65 }
  b := Person{ "Barack Obama", 54 }
  fmt.Println(n, b)
}
