// chapter 1, exercise 2

// Modify the echo program to print the index and value of the each of its
// arguments one per line

package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[:] {
		fmt.Println(index, arg)
	}
}
