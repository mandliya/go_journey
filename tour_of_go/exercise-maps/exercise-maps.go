/**
 *  maps exercise
 *  Implement WordCount. It should return a map of the counts of each “word” in the string s.
 *  The wc.Test function runs a test suite against the provided function and prints success or failure.
 */

package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  var wordmap map[string]int
  wordmap = make(map[string]int)
  words := strings.Fields(s)
  for _,word := range words {
    wordmap[word] += 1
  }
  return wordmap
}


func main() {
  wc.Test(WordCount)
}
