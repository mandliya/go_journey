/**
 * interface in go are satisfied implicitly
 * A type implements an interface by implementing the methods.
 * There is no explicit declaration of intent; no "implements" keyword.
 *
 * Implicit interfaces decouple implementation packages from the packages that define the interfaces:
 *  neither depends on the other.
 *  It also encourages the definition of precise interfaces,
 *  because you don't have to find every implementation and tag it with the new interface name.
 *  Here io already defines Read and Write methods.
 */

 package main

 import (
   "fmt"
   "os"
 )

 type Writer interface {
   Write(b []byte) (n int, err error)
 }

 type Reader interface {
   Read(b []byte) (n int, err error)
 }

 type ReaderWriter interface {
  Reader
  Writer
}

func main() {
  var w Writer
  w = os.Stdout
  fmt.Fprintf(w, "Hello Writer\n")
}
