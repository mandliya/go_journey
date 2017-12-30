// Chapter 1, exercise 4
// Modify dup2 to print the names of all the files in which duplicate lines occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

// LineInfo Represents a line's occurance count and the files it belong to
type LineInfo struct {
	Count     int
	Filenames []string
}

func main() {
	counts := make(map[string]*LineInfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex4: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, lineInfo := range counts {
		if lineInfo.Count > 1 {
			fmt.Printf("%d\t%v\n%s\n", lineInfo.Count, lineInfo.Filenames, line)
		}
	}
}

func countLines(f *os.File, counts map[string]*LineInfo) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		_, ok := counts[key]
		if ok {
			counts[key].Count++
			counts[key].Filenames = append(counts[key].Filenames, f.Name())
		} else {
			counts[key] = new(LineInfo)
			counts[key].Count = 1
			counts[key].Filenames = append(counts[key].Filenames, f.Name())
		}
	}
}
