/**
 * The range form of the for loop iterates over a slice or map.
 * You can skip the index or value by assigning to _.
 * If you only want the index, drop the ", value" entirely.
 */

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256}

func main() {
  for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
  }

  //skipping index
  for _,value := range pow {
    fmt.Printf("%d, ", value)
  }
  fmt.Printf("\n")
  //skipping value
  for index := range pow {
    fmt.Printf("%d, ", index)  
  }
  fmt.Printf("\n")
}
