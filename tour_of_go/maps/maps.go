/**
 * A map maps keys to values.
 * NOTE Maps must be created with make (not new) before use;
 * the nil map is empty and cannot be assigned to.
 */

package main

import "fmt"

type Location struct {
  Long, Lat float64
}

func main() {
  var m map[string]Location
  //NOTE make it before using it
  m = make(map[string]Location)
  m["Google"] = Location{37.422319, -122.084015}
  fmt.Println("Latitude and Longitude of Google Headquaters are ", m["Google"])
}
