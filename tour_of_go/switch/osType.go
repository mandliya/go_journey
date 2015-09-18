/**
 *  Switch statement in Go, almost similar to c/c++
 *  As we saw in if else, we will use short statement here to determine val of os
 *  and use it in switch scope
 */

package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("Mac OS X")
  case "linux" :
    fmt.Println("Linux")
  default:
    fmt.Println("%s\n", os)
  }
}
