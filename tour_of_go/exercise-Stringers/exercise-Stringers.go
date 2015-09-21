/**
 * Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.
 * For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
 */

package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string{
  return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
  addrs := map[string]IPAddr {
    "LoopBack"  : { 127, 0, 0, 1},
    "GoogleDNS" : { 8, 8, 8, 8},
  }
  for n,a := range addrs {
    fmt.Printf("%v : %v\n", n, a)
  }
}
