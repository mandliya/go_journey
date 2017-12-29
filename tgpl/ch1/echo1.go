// Naive implementation of Unix like Echo.
// This program shows how to print command line arguments.
// We use OS package of Go.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
