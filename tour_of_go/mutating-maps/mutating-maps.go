/**
 *  Insert, delete, retrieve an element in map
 */

package main

import "fmt"

func main() {
  var answerMap map[string]int
  answerMap = make(map[string]int)

  answerMap["problem1"] = 42
  fmt.Println("The answer of problem 1 is ", answerMap["problem1"])

  //updating answer
  answerMap["problem1"] = 21
  fmt.Println("The updated answer of problem 1 is ", answerMap["problem1"])

  //deleting answer
  delete(answerMap, "problem1")
  fmt.Println("The answer of problem 1 is removed ", answerMap["problem1"])

  //checking if answer exists
  v, ok := answerMap["problem1"]
  fmt.Println("The value ", v, "Present?", ok)
}
