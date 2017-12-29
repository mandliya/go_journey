// Chapter 1, exercise 3
// Experiment to measure the the difference in running time between our
// potentially inefficient versions and the one that uses strings.Join.
//

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// version 1:
	start := time.Now()
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	end := time.Now()
	fmt.Println(end.Sub(start))

	// version 2
	start = time.Now()
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	end = time.Now()
	fmt.Println(end.Sub(start))

	// version 3
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end = time.Now()
	fmt.Println(end.Sub(start))
}
