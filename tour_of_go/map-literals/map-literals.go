/**
 * Map literals are like struct literals, but the keys are required.
 */

package main

import "fmt"

type Location struct {
  Long, Lat float64
}

func main() {
  var wonders = map[string]Location {
    "The Great Wall of China" : Location{40.437483, 116.571664},
    "The Taj Mahal"           : Location{27.175010, 78.042209},
    "The Eiffel Tower"        : Location{48.858578, 2.294492},
  }
  fmt.Println(wonders)

  //If top level type is just a name you can omit it from the elments of the literals
  // For example
  var cities = map[string]Location {
    "New Delhi" : {28.617115, 77.244566},
    "New York"  : {40.713202, -74.004400},
    "Paris"     : {48.856713, 2.353393},
  }
  fmt.Println(cities)
}
